/**
* @file: server.go ==> gateway/router
* @package: router
* @author: jingxiu
* @since: 2022/11/8
* @desc: 网络服务入口
 */

package router

import (
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	//设置gin模式
	//gin.SetMode(global.VP.GetString("RunMode"))
	engine := gin.New()
	// 使用日志打印
	engine.Use(gin.Logger())
	//定义404中间件
	engine.NoRoute(middleware.NoRoute())
	//	 路由挂载
	loadRouter(engine)
	return engine
}

func loadRouter(engine *gin.Engine) {
	// TODO: add your router group...
}
