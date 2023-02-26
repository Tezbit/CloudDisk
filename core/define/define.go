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
var AuthPwd = "SILFWOFGDWOPMLEG"

//验证码过期时间(s)
var CondeExpire = 300
