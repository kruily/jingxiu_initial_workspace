/**
* @file: router.go ==> gateway/router
* @package: router
* @author: jingxiu
* @since: 2022/12/3
* @desc: //TODO
 */

package router

import "github.com/gin-gonic/gin"

type (
	Route interface {
		Route(*gin.Engine, func(*gin.Engine))
	}
	RegisterRoute struct {
		Ro Route
		Do func(*gin.Engine)
	}
)

var registers []RegisterRoute

func loadRouter(engine *gin.Engine) {
	for _, re := range registers {
		re.Ro.Route(engine, re.Do)
	}
}

func Register(ro ...RegisterRoute) {
	registers = append(registers, ro...)
}
