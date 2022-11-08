/**
* @file: svc.go ==> gateway/services
* @package: services
* @author: jingxiu
* @since: 2022/11/8
* @desc: 微服务链接定义
 */

package services

import (
	"gateway/config"
	"go.uber.org/zap"
)

var SvcContext *ServiceContext

type ServiceContext struct {
	Grpc *GrpcContext
	Log  *zap.Logger
}

func NewContext(c *config.Config) *ServiceContext {
	grpc := GrpcInit(c)
	return &ServiceContext{
		Grpc: grpc,
		Log:  zap.NewExample(),
	}
}
