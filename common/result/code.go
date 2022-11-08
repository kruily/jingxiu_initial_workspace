/**
* @file: code.go ==> common/result
* @package: result
* @author: jingxiu
* @since: 2022/11/8
* @desc: 错误码定义
*	错误码需为 > 0 的数，反之表示正确。
*	服务级别错误码：1 位数进行表示，比如 1 为系统级错误；2 为普通错误，通常是由用户非法操作引起。
*	模块级错误码：2 位数进行表示，比如 01 为用户模块；02 为订单模块。
*	具体错误码：2 位数进行表示，比如 01 为手机号不合法；02 为验证码输入错误。
 */

package result

var (
	// OK
	OK = NewResult(0, "OK")

	// 服务级错误码
	ErrServer    = NewResult(10001, "服务异常，请联系管理员")
	ErrNetwork   = NewResult(10002, "服务器开小差了...")
	ErrParam     = NewResult(10003, "参数有误")
	ErrSignParam = NewResult(10004, "签名参数有误")
	ErrNoRoute   = NewResult(404, "路由不存在")

	// token校验错误
	ErrToken       = NewResult(10005, "登录失效")
	ErrNoUser      = NewResult(10006, "用户未登录")
	ErrUserIllegal = NewResult(10007, "非法用户")

	// 模块级错误码 - 应用模块
	// ....

)
