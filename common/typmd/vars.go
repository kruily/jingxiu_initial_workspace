/**
* @file: vars.go ==> common/typmd
* @package: typmd
* @author: jingxiu
* @since: 2022/12/1
* @desc: //TODO
 */

package typmd

import "errors"

var (
	LinkedIsEmptyError     = errors.New("the linked list is empty")
	LinkedCallBackError    = errors.New("the callback function failed to execute")
	LinkedNoneNodeError    = errors.New("there are no elements")
	LinkedInvalidNodeError = errors.New("invalid node data")
)
