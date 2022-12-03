/**
* @file: singleLinked_test.go ==> common/typmd
* @package: typmd
* @author: jingxiu
* @since: 2022/12/1
* @desc: //TODO
 */

package typmd

import (
	"common/utils"
	"fmt"
	"testing"
)

type T struct {
	Name string
}

func TestSingleLinked_NewSingleLinked(t *testing.T) {
	t1 := NewSingleLinked()
	for i := 0; i < 100; i++ {
		if err := t1.Insert(&T{Name: "t" + utils.Itoa(i)}); err != nil {
			fmt.Printf(err.Error())
			return
		}
	}
	f := t1.Find(func(item *SingleLinked) bool {
		if item.Extend.(*T).Name == "t3" {
			return true
		} else {
			return false
		}
	})
	fmt.Printf("%#v\n", f)
	f, _ = t1.Update(func(item *SingleLinked) bool {
		if item.Extend.(*T).Name == "t3" {
			return true
		} else {
			return false
		}
	}, func(item *SingleLinked) error {
		item.Extend.(*T).Name = "T3"
		return nil
	})
	fmt.Printf("%#v\n", f.Extend)
	if err := t1.Remove(func(item *SingleLinked) bool {
		if item.Extend.(*T).Name == "T3" {
			return true
		} else {
			return false
		}
	}); err != nil {
		fmt.Println(err.Error())
	}
	temp := t1.Next
	for temp != nil {
		fmt.Printf("%d\t%#v\n", temp.Pos, temp.Extend)
		temp = temp.Next
	}
}
