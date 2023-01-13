/**
* @file: middleware.go ==> gateway/middleware
* @package: middleware
* @author: jingxiu
* @since: 2023/1/13
* @desc: //TODO
 */

package middleware

import (
	"github.com/gin-gonic/gin"
)

func init() {
	sort()
}

type Middleware struct {
	Func func() gin.HandlerFunc // 中间件
	Use  bool                   // 是否使用
	Sort int                    // 排序
}

var registerList []Middleware

func register(f func() gin.HandlerFunc, b bool, s int) {
	registerList = append(registerList, Middleware{
		Func: f,
		Use:  b,
		Sort: s,
	})
}

// 希尔排序
func sort() {
	length := len(registerList)
	if length <= 1 {
		return
	}
	step := length / 2
	for ; step > 0; step = step / 2 {
		for i := step; i < length; i++ {
			j := i - step
			temp := registerList[i]
			for j >= 0 && temp.Sort > registerList[j].Sort {
				registerList[j+step] = registerList[j]
				j = j - step
			}
			registerList[j+step] = temp
		}
	}
}

func UseMiddle(e *gin.Engine) {
	for _, v := range registerList {
		if v.Use {
			e.Use(v.Func())
		}
	}
}
