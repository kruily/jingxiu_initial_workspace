/**
* @file: svc.go ==> gateway/services
* @package: services
* @author: jingxiu
* @since: 2022/11/8
* @desc: 微服务链接定义
 */

package services

import (
	"common/log"
	"gateway/config"
	"github.com/sirupsen/logrus"
)

var SvcContext *ServiceContext

type ServiceContext struct {
	Grpc *GrpcContext
	Log  *logrus.Entry
}

func NewContext(c *config.Config) *ServiceContext {
	grpc := GrpcInit(c)
	return &ServiceContext{
		Grpc: grpc,
		Log:  log.GetLogEntry(c.Logger.FilePath, c.Logger.FileName),
	}
}
