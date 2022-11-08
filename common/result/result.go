/**
* @file: result.go ==> common/result
* @package: result
* @author: jingxiu
* @since: 2022/11/8
* @desc: 请求返回体系
 */

package result

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

var _ Result = (*res)(nil)

type Result interface {
	// i 为了避免被其他包实现
	i()
	// WithData 设置成功时返回的数据
	WithData(data interface{}) Result
	// WithID 设置当前请求的唯一ID
	WithID(id string) Result
	// WithMsg 设置当前返回的消息信息
	WithMsg(msg string) Result
	// ToString 返回 JSON 格式的错误详情
	ToString() string
	// Response	gin的消息返回
	Response(*gin.Context)
}

type res struct {
	Code int         `json:"code"`         // 业务编码
	Msg  string      `json:"msg"`          // 错误描述
	Data interface{} `json:"data"`         // 成功时返回的数据
	ID   string      `json:"id,omitempty"` // 当前请求的唯一ID，便于问题定位，忽略也可以
}

func (e *res) i() {}

func (e *res) WithData(data interface{}) Result {
	e.Data = data
	return e
}

func (e *res) WithID(id string) Result {
	e.ID = id
	return e
}
func (e *res) WithMsg(msg string) Result {
	e.Msg = msg
	return e
}

// ToString 返回 JSON 格式的错误详情
func (e *res) ToString() string {
	err := &struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
		ID   string      `json:"id,omitempty"`
	}{
		Code: e.Code,
		Msg:  e.Msg,
		Data: e.Data,
		ID:   e.ID,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}

func (e *res) Response(c *gin.Context) {
	c.JSONP(e.Code, e)
}

func NewResult(code int, msg string) Result {
	return &res{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
