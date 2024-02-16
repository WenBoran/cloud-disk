package define

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.RegisteredClaims
}

var (
	JwtKey       = "cloud-disk"
	MailPassword = "5j7UqvoichJnjiGc"
)

// 验证码长度
var CodeLength = 6

// 验证码过期时间
var CodeExpire = 2 * time.Minute
