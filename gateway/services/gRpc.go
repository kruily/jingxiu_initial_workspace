/**
* @file: gRpc.go ==> gateway/services
* @package: services
* @author: jingxiu
* @since: 2022/11/8
* @desc:gRPC 连接器
 */

package services

import "gateway/config"

// GrpcContext 填入gRpc客户端注册实例
type GrpcContext struct {
	// 示例 添加用户服务rpc
	// UserRpc     user.UserClient
}

// GrpcInit 初始化gRpc链接
func GrpcInit(c *config.Config) *GrpcContext {
	var g GrpcContext
	// 示例 添加用户服务rpc
	//user_conn, err := grpc.Dial(c.UserRpc.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	panic(err)
	//}
	//g.UserRpc = user.NewUserClient(user_conn)

	return &g
}
