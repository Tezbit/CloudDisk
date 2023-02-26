package helper

import (
	"cloud_disk/core/define"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/smtp"
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

func GenerateToken(id uint64, identity, name string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func MailSendCode(to, code string) error {
	e := email.NewEmail()
	e.From = "Get <guzxrag61971@163.com>"
	e.To = []string{to}
	e.Subject = "CloudDisk验证码"
	e.HTML = []byte("你的验证码为：<h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "guzxrag61971@163.com", define.AuthPwd, "smtp.163.com"),
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
