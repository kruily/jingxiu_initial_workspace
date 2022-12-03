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
	"gateway/services"
)

func main() {
	//1. read config from yaml
	if err := config.ConfigInit("./gateway.yaml"); err != nil {
		panic("config file read ersor!")
		return
	}

	// 2. start some services
	services.SvcContext = services.NewContext(config.C)

	// 3. start gin router
	r := router.GinApplication()
	router.Server(r)
	//if err := r.Run(config.C.Gateway.Listen); err != nil {
	//	fmt.Printf("startup service failed, err:%v\n\n", err)
	//}
}
