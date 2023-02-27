package test

import (
	"bytes"
	"cloud_disk/core/define"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestFileUploadByFilepath(t *testing.T) {
	u, _ := url.Parse(define.Bucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.tencentcloud.com/cam/capi
			SecretID: os.Getenv("TencentSecretID"), // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://www.tencentcloud.com/document/product/598/37140?from_cn_redirect=1
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.tencentcloud.com/cam/capi
			SecretKey: os.Getenv("TencentSecretKey"), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://www.tencentcloud.com/document/product/598/37140?from_cn_redirect=1
		},
	})
	key := "cloud-disk/README.md"

	_, _, err := client.Object.Upload(
		context.Background(), key, "../README.md", nil,
	)
	if err != nil {
		panic(err)
	}
}

func TestFileUploadByReader(t *testing.T) {
	u, _ := url.Parse(define.Bucket)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.tencentcloud.com/cam/capi
			SecretID: os.Getenv("TencentSecretID"), // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://www.tencentcloud.com/document/product/598/37140?from_cn_redirect=1
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.tencentcloud.com/cam/capi
			SecretKey: os.Getenv("TencentSecretKey"), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://www.tencentcloud.com/document/product/598/37140?from_cn_redirect=1
		},
	})

	key := ""

	f, err := os.ReadFile("")
	if err != nil {
		return
	}
	_, err = client.Object.Put(
		context.Background(), key, bytes.NewReader(f), nil,
	)
	if err != nil {
		panic(err)
	}
}
