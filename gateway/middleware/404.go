/**
* @file: 404.go ==> gateway/middleware
* @package: middleware
* @author: jingxiu
* @since: 2022/11/2
* @desc: 当出现不存的路由时，返回404
 */

package middleware

import (
	"common/result"
	"github.com/gin-gonic/gin"
)

func NoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 抛出 404 路由不存在
		result.ErrNoRoute.Response(c)
	}
}
