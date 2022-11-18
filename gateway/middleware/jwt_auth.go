/**
* @file: jwt_auth.go ==> gateway/middleware
* @package: middleware
* @author: jingxiu
* @since: 2022/11/4
* @desc: //TODO
 */

package middleware

import (
	"common/jwt"
	"common/result"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("authorization")
		if authHeader == "" {
			//c.JSON(http.StatusUnauthorized, gin.H{
			//	"code": 2003,
			//	"msg":  "用户未登录",
			//})
			result.ErrToken.Response(c)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.Split(authHeader, ".")
		if len(parts) != 3 {
			result.ErrToken.WithMsg("非法token").Response(c)
			c.Abort()
			return
		}
		mc, ok := jwt.ParseToken(authHeader)
		if ok != nil {
			result.ErrToken.WithMsg("非法token").Response(c)
			c.Abort()
			return
		}
		// 即将超过过期时间，则刷新token 暂不需要
		//if mc.ExpiresAt < time.Now().Unix()+300 {
		//	err := common.RefreshToken(c, mc)
		//	if err != nil {
		//		result.ErrToken.WithMsg("登录已失效，请重新登录").Response(c)
		//		c.Abort()
		//		return
		//	}
		//}
		c.Set("user", mc.Jwtkey)
		c.Set("Expires", mc.ExpiresAt)
		c.Next()
	}
}
