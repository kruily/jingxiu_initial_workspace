/**
* @file: main.go ==> gateway
* @package: gateway
* @author: jingxiu
* @since: 2022/11/8
* @desc: 程序主入口
 */

package main

import (
	"flag"
	"gateway/config"
	"gateway/router"
)

var gatewayConfig = flag.String("f", "../etc/gateway.yaml", "网关层配置文件路径")

// @title Swagger API Docs
// @version
// @description
// @termsOfService
// @contact.name
// @contact.url
// @contact.email
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8000
// @BasePath /
func main() {
	flag.Parse()
	// read config from yaml
	if err := config.ConfigInit(*gatewayConfig); err != nil {
		panic("config file read error!")
		return
	}

	// start some rpc services
	//services.SvcContext = services.NewContext(config.C)

	// start gin router
	r := router.GinApplication()
	router.Server(r)
}
