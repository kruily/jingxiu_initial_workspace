/**
* @file: int.go ==> common/typmd
* @package: typmd
* @author: jingxiu
* @since: 2022/12/1
* @desc: //TODO
 */

package typmd

type (
	Integer   int
	Integer32 int32
	Integer64 int64
)

// Compare Integer类型比较
func (i Integer) Compare(d int) bool {
	return i > Integer(d)
}

// Compare Integer类型比较
func (i Integer32) Compare(d int) bool {
	return i > Integer32(d)
}

// Compare Integer类型比较
func (i Integer64) Compare(d int) bool {
	return i > Integer64(d)
}

// Equal Integer类型比较相等
func (i Integer) Equal(d int) bool {
	return i == Integer(d)
}

// Equal Integer类型比较相等
func (i Integer32) Equal(d int) bool {
	return i == Integer32(d)
}

// Equal Integer类型比较相等
func (i Integer64) Equal(d int) bool {
	return i == Integer64(d)
}

// Add Integer类型 值相加
func (i Integer) Add(d int) Integer {
	i += Integer(d)
	return i
}

// Add Integer类型 值相加
func (i Integer32) Add(d int) Integer32 {
	i += Integer32(d)
	return i
}

// Add Integer类型 值相加
func (i Integer64) Add(d int) Integer64 {
	i += Integer64(d)
	return i
}

// Sub Integer类型 值相减
func (i Integer) Sub(d int) Integer {
	i -= Integer(d)
	return i
}

// Sub Integer类型 值相减
func (i Integer32) Sub(d int) Integer32 {
	i -= Integer32(d)
	return i
}

// Sub Integer类型 值相减
func (i Integer64) Sub(d int) Integer64 {
	i -= Integer64(d)
	return i
}
