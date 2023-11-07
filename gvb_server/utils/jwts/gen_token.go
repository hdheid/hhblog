package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
	"gvb_server/global"
	"time"
)

// GenToken 创建token
func GenToken(user JwtPayLoad) (string, error) {
	MySecret = []byte(global.Config.Jwy.Secret)
	Issuer := global.Config.Jwy.Issuer
	ExpirationTime := global.Config.Jwy.Expires

	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(ExpirationTime))), // 过期时间
			Issuer:    Issuer,                                                            // 签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}
