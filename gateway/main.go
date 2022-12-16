/**
* @file: main.go ==> gateway
* @package: gateway
* @author: jingxiu
* @since: 2022/11/8
* @desc: 程序主入口
 */

package main

import (
	"gateway/config"
	"gateway/router"
)

func main() {
	// read config from yaml
	if err := config.ConfigInit("./gateway.yaml"); err != nil {
		panic("config file read error!")
		return
	}

	// start some services
	//services.SvcContext = services.NewContext(config.C)

	// start gin router
	r := router.GinApplication()
	router.Server(r)
}
