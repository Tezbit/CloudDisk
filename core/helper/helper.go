package helper

import (
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
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"path"
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
