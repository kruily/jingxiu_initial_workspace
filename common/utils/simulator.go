/**
* @file: simulator.go ==> common/utils
* @package: utils
* @author: jingxiu
* @since: 2022/11/8
* @desc: 一些模拟运算
 */

package utils

// ThreeUnaryInterface 三元模拟器 抽象类型模式
func ThreeUnaryInterface(condition bool, t, f interface{}) interface{} {
	if condition {
		return t
	}
	return f
}

// ThreeUnaryFunction 三元模拟器 方法模式
func ThreeUnaryFunction(condition bool, t, f func()) {
	if condition {
		t()
	}
	f()
}
