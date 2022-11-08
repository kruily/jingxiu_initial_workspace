/**
* @file: slice.go ==> common/format
* @package: format
* @author: jingxiu
* @since: 2022/11/8
* @desc: 切片转化
 */

package format

import "fmt"

// Reverse 反转切片
func Reverse[T string | int | int32 | int64](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// SliceToString 将一个切片数据，转换成一个元组包含的字符串
func SliceToString[T int64 | int32 | string | float64](slice []T) string {
	s := "("
	for _, item := range slice {
		ns := fmt.Sprintf("%v", item)
		s = s + ns + ","
	}
	s = s[:len(s)-1]
	s = s + ")"
	return s
}
