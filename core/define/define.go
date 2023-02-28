package define

import "github.com/dgrijalva/jwt-go"

/*
 * 默认的变量设置
 */

type UserClaim struct {
	Id       uint64
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-tezbit"

//验证码过期时间(s)
var CondeExpire = 300

var TokenExpire = 3600
var RefreshTokenExpire = 7200

var Bucket = "https://tezbit-1316751748.cos.ap-nanjing.myqcloud.com"

//分页默认参数
var PageSize = 20

var Datetime = "2006-01-02 15:04:05"
