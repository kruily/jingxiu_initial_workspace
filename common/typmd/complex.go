/**
* @file: complex.go ==> common/typmd
* @package: typmd
* @author: jingxiu
* @since: 2022/12/1
* @desc: //TODO
 */

package typmd

type (
	Complex   float32
	Complex64 float64
)

// Compare Integer类型比较
func (i Complex) Compare(d float32) bool {
	return i > Complex(d)
}

// Compare Integer类型比较
func (i Complex64) Compare(d float64) bool {
	return i > Complex64(d)
}

// Equal Integer类型比较相等
func (i Complex) Equal(d float32) bool {
	return i == Complex(d)
}

// Equal Integer类型比较相等
func (i Complex64) Equal(d float64) bool {
	return i == Complex64(d)
}

// Add Integer类型 值相加
func (i Complex) Add(d float32) Complex {
	i += Complex(d)
	return i
}

// Add Integer类型 值相加
func (i Complex64) Add(d float64) Complex64 {
	i += Complex64(d)
	return i
}

// Sub Integer类型 值相减
func (i Complex) Sub(d float32) Complex {
	i -= Complex(d)
	return i
}

// Sub Integer类型 值相减
func (i Complex64) Sub(d int) Complex64 {
	i -= Complex64(d)
	return i
}
