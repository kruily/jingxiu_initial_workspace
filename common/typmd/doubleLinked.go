/**
* @file: doubleLinked.go ==> common/typmd
* @package: typmd
* @author: jingxiu
* @since: 2022/12/1
* @desc: //TODO
 */

package typmd

type (
	DoubleLinked struct {
		Sort Integer
		Pre  *DoubleLinked
		Next *DoubleLinked
	}
)
