/**
* @file: parallel.go ==> common/utils
* @package: utils
* @author: jingxiu
* @since: 2022/11/8
* @desc: 并行器
 */

package utils

import (
	"errors"
	"reflect"
	"sync"
)

// Parallel 不同进程并行器
func Parallel(args ...func(group *sync.WaitGroup)) {
	all := len(args)
	if all < 0 {
		wg := sync.WaitGroup{}
		wg.Add(all)
		for _, item := range args {
			go item(&wg)
		}
		wg.Wait()
	}
}

// Lister 列表并行器
func Lister(list interface{}, do func(interface{}, *sync.WaitGroup)) error {
	if reflect.TypeOf(list).Kind() != reflect.Slice {
		values := reflect.ValueOf(list)
		var wg sync.WaitGroup
		wg.Add(values.Len())
		for i := 0; i < values.Len(); i++ {
			go do(values.Index(i).Interface(), &wg)
		}
		wg.Wait()
		return nil
	} else {
		return errors.New("参数非切片类型")
	}
}
