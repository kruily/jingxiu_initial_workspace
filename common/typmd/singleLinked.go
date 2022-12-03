/**
* @file: singleLinked.go ==> common/typmd
* @package: typmd
* @author: jingxiu
* @since: 2022/12/1
* @desc: //TODO
 */

package typmd

import "fmt"

type (
	SingleLinked struct {
		Pos    int           // 节点所处位置
		Next   *SingleLinked // 下个节点
		Extend any
	}

	LinkedIF interface {
		Insert(any) error                                // 头部插入
		Remove(LinkedCallBack) error                     // 移除
		Clear()                                          // 清空
		Find(LinkedCallBackBool) error                   // 查询
		Update(LinkedCallBackBool, LinkedCallBack) error // 更新
		Reverse() error                                  // 反转
		IsEmpty() bool                                   // 是否为空
		Wrap(SingleLinked) error                         // 交换
		Equal(*SingleLinked, LinkedCallBackBoolTow) bool // 比较相等
		Sort(LinkedCallBackBoolTow)                      // 排序
	}

	LinkedCallBack        func(*SingleLinked) error
	LinkedCallBackBool    func(*SingleLinked) bool
	LinkedCallBackBoolTow func(*SingleLinked, *SingleLinked) bool
)

// NewSingleLinked 初始化一个链表
func NewSingleLinked() *SingleLinked {
	return &SingleLinked{
		Pos:  -1,
		Next: nil,
	}
}

// IsEmpty 判断链表是否为空
func (l *SingleLinked) IsEmpty() bool {
	if l == nil || l.Next == nil {
		return true
	}
	return false
}

// Clear 清空链表
func (l *SingleLinked) Clear() {
	l.Next = nil
	l.Pos = 0
	l.Next = nil
}

// Insert 单向链表头部插入
func (l *SingleLinked) Insert(node any) error {
	if l.Next == nil {
		l.Next = &SingleLinked{
			Pos:    0,
			Extend: node,
			Next:   nil,
		}
	} else {
		next := l.Next
		l.Next = &SingleLinked{
			Pos:    next.Pos + 1,
			Extend: node,
			Next:   next,
		}
	}
	return nil
}

// Remove 单向链表移除某个指定的元素
// Parameter do:LinkedCallBackBool 用于自定义指定元素的回调函数
func (l *SingleLinked) Remove(do LinkedCallBackBool) error {
	temp := l
	for temp != nil {
		if temp.Next.Extend == nil {
			return LinkedInvalidNodeError
		}
		if do(temp.Next) {
			fmt.Println("true")
			temp.Next = temp.Next.Next
			return nil
		}
		temp = temp.Next
	}
	return LinkedNoneNodeError
}

// Find 查询某个指定元素
// Parameter do:LinkedCallBackBool 用于自定义寻找的元素标准的回调函数
func (l *SingleLinked) Find(do LinkedCallBackBool) *SingleLinked {
	temp := l.Next
	for temp != nil {
		if do(temp) {
			return temp
		}
		temp = temp.Next
	}
	return nil
}

// Update 更新指定的元素
// Parameter do:LinkedCallBackBool 自定义查询元素方式的回调函数；up:LinkedCallBack 自定以更新元素方式的回调函数
func (l *SingleLinked) Update(do LinkedCallBackBool, up LinkedCallBack) (*SingleLinked, error) {
	f := l.Find(do)
	if err := up(f); err != nil {
		return nil, err
	}
	return f, nil
}

// Reverse 反转链表
func (l *SingleLinked) Reverse() error {
	if l == nil {
		return LinkedIsEmptyError
	}
	return nil
}

// Wrap 链表两节点交换
// Parameter one,other:SingleLinked 传递两个节点
func (l *SingleLinked) Wrap(one, other *SingleLinked) error {
	if l == nil {
		return LinkedIsEmptyError
	}
	temp := one.Extend
	one.Extend = other.Extend
	other.Extend = temp
	return nil
}

// Equal 比较两节点是否相等
// Parameter one,other:SingleLinked 传递两个节点；do:LinkedCallBackBoolTow 比较方式
func (l *SingleLinked) Equal(one, other *SingleLinked, do LinkedCallBackBoolTow) bool {
	return do(one, other)
}

// Sort 链表根据某种方式排序
// Parameter
func (l *SingleLinked) Sort(do LinkedCallBackBoolTow) error {
	if l == nil {
		return LinkedIsEmptyError
	}
	i := l.Next
	j := i.Next
	for i != nil {
		for i.Next != nil || j.Next != nil {
			if do(i.Next, j.Next) {
				if err := l.Wrap(i, j); err != nil {
					return err
				}
			}
		}
		i = i.Next
		j = j.Next
	}
	return nil
}
