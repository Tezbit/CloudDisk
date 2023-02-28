package helper

import (
	"bytes"
	"cloud_disk/core/define"
	"cloud_disk/core/internal/svc"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"golang.org/x/crypto/bcrypt"
	"io"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

func Generate(s string) (string, error) {
	byt, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	return string(byt), err
}

func Validate(s string, hashed string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(s)); err != nil {
		return false, err
	}
	return true, nil
}

func GenerateToken(id uint64, identity, name string, sec int64) (string, error) {
	uc := define.UserClaim{
		Id:             id,
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Second * time.Duration(sec)).Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func MailSendCode(from, to, code, pwd string) error {
	e := email.NewEmail()
	e.From = "Get <" + from + ">"
	e.To = []string{to}
	e.Subject = "CloudDisk验证码"
	e.HTML = []byte("你的验证码为：<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", from, pwd, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	return err
}

//6位验证码生成
func Captcha() string {
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		code = fmt.Sprintf("%s%d", code, rand.Intn(10))
	}
	return code
}

func GetUUid() string {
	return uuid.NewV4().String()
}

func CosUpload(ctx *svc.ServiceContext, r *http.Request) (string, error) {
	//todo :复用部分抽出来
	u, _ := url.Parse(define.Bucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.tencentcloud.com/cam/capi
			SecretID: ctx.CosSecret.SecretID, // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://www.tencentcloud.com/document/product/598/37140?from_cn_redirect=1
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.tencentcloud.com/cam/capi
			SecretKey: ctx.CosSecret.SecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://www.tencentcloud.com/document/product/598/37140?from_cn_redirect=1
		},
	})

	file, fileHeader, err := r.FormFile("file")
	key := "cloud-disk/" + GetUUid() + path.Ext(fileHeader.Filename)

	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)
	if err != nil {
		panic(err)
	}
	return define.Bucket + "/" + key, nil
}

// Token 解析
func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}

// CosInitPart 分片上传初始化
func CosInitPart(ctx *svc.ServiceContext, ext string) (string, string, error) {
	//todo :复用部分抽出来
	u, _ := url.Parse(define.Bucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  ctx.CosSecret.SecretID,
			SecretKey: ctx.CosSecret.SecretKey,
		},
	})
	key := "cloud-disk/" + GetUUid() + ext
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		return "", "", err
	}
	return key, v.UploadID, nil
}

// CosPartUpload 分片上传
func CosPartUpload(ctx *svc.ServiceContext, r *http.Request) (string, error) {
	//todo :复用部分抽出来
	u, _ := url.Parse(define.Bucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  ctx.CosSecret.SecretID,
			SecretKey: ctx.CosSecret.SecretKey,
		},
	})
	key := r.PostForm.Get("key")
	UploadID := r.PostForm.Get("upload_id")
	partNumber, err := strconv.Atoi(r.PostForm.Get("part_number"))
	if err != nil {
		return "", err
	}
	f, _, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, f)

	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, partNumber, bytes.NewReader(buf.Bytes()), nil,
	)
	if err != nil {
		return "", err
	}
	return strings.Trim(resp.Header.Get("ETag"), "\""), nil
}

// CosPartUploadComplete 分片上传完成
func CosPartUploadComplete(ctx *svc.ServiceContext, key, uploadId string, co []cos.Object) error {
	//todo :复用部分抽出来
	u, _ := url.Parse(define.Bucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  ctx.CosSecret.SecretID,
			SecretKey: ctx.CosSecret.SecretKey,
		},
	})

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, co...)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, uploadId, opt,
	)
	return err
}
