/**
* @file: jwt.go ==> common/jwt
* @package: jwt
* @author: jingxiu
* @since: 2022/11/8
* @desc: jwt 生成
 */

package jwt

import (
	"fmt"
	"gateway/config"
	"github.com/gin-gonic/gin"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Jwtkey string `json:"Jwtkey"`
	// 可自行添加选择参数
	jwt.StandardClaims
}

func GenerateToken(Jwtkey string, jwtSecret []byte) (string, error) {
	num := time.Duration(config.C.Jwt.Expire)
	var mp = map[string]time.Duration{
		"s": num * time.Second,
		"m": num * time.Minute,
		"h": num * time.Hour,
	}
	t := time.Now().Add(mp[config.C.Jwt.Unit])
	claims := Claims{
		Jwtkey,
		jwt.StandardClaims{
			ExpiresAt: t.Unix(),
			Issuer:    "jingxiu@center",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.C.Jwt.Secret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			fmt.Println(claims)
			return claims, nil
		}
	}
	return nil, err
}

func RefreshToken(c *gin.Context, mc *Claims) error {
	if mc.ExpiresAt == 0 {
		expires, _ := c.Get("Expires")
		mc.ExpiresAt = expires.(int64)
	}
	token, err := GenerateToken(mc.Jwtkey, []byte(config.C.Jwt.Secret))
	if err != nil {
		return err
	}
	c.Header("refresh", token)
	return nil
}
