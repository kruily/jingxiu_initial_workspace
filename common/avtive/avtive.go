/**
* @file: avtive.go ==> common
* @package: common
* @author: jingxiu
* @since: 2022/11/8
* @desc: 一个简述用户活跃信息的包
 */

package avtive

import (
	"fmt"
	"time"
)

var timeTemplates = []string{
	// "2006-01-02 15:04:05", //常规类型
	// "2006/01/02 15:04:05",
	// "2006-01-02",
	// "2006/01/02",
	// "15:04:05",
	"2006-01-02T15:04:05+08:00",
	"2006-01-02 15:04:05.1051598 +0800 CST",
	"2006-01-02 15:04:05 +0800 CST",
}

// Active 定义活跃状态结构体
type Active struct {
	Minutes       int64     `json:"minutes"`
	Hours         int64     `json:"hours"`
	Days          int64     `json:"days"`
	Months        int64     `json:"months"`
	Years         int64     `json:"years"`
	Description   string    `json:"description"`
	LastTimestamp time.Time `json:"last_timestamp"`
	NowTimestamp  time.Time `json:"now_timestamp"`
}

func New(last string) *Active {
	// 将string转化为time.Time类型
	ret := &Active{
		LastTimestamp: TimeStringToGoTime(last),
		NowTimestamp:  time.Now(),
	}
	// 通过差值计算Active其他参数
	duration := ret.GetDuration()
	ret.Minutes = duration / 60
	ret.Hours = ret.Minutes / 60
	ret.Days = ret.Hours / 24
	ret.Months = ret.Days / 30
	ret.Years = ret.Months / 12
	ret.GenerateDesc()
	return ret
}

func TimeStringToGoTime(tm string) time.Time {
	for i := range timeTemplates {
		t, err := time.ParseInLocation(timeTemplates[i], tm, time.Local)
		fmt.Printf("%#v\n", t)
		if nil == err && !t.IsZero() {
			return t
		}
	}

	return time.Time{}
}

// GetDuration 获取两个时间戳的差值
func (a *Active) GetDuration() int64 {
	return a.NowTimestamp.Unix() - a.LastTimestamp.Unix()
}

// GenerateDesc 根据分钟生成描述
func (a *Active) GenerateDesc() {
	if a.Minutes >= 60 {
		// 去计算小时描述生成
		a.HoursDesc()
	} else if a.Minutes >= 0 && a.Minutes < 5 {
		a.Description = "刚刚活跃"
	} else if a.Minutes >= 5 && a.Minutes < 10 {
		a.Description = "5分钟前活跃"
	} else if a.Minutes >= 10 && a.Minutes < 30 {
		a.Description = "10分钟前活跃"
	} else if a.Minutes >= 30 && a.Minutes < 60 {
		a.Description = "30分钟前活跃"
	}
}

// HoursDesc 根据小时生成 描述
func (a *Active) HoursDesc() {
	if a.Hours > 24 {
		// 去计算天生成描述
		a.DaysDesc()
	} else {
		a.Description = fmt.Sprintf("%v小时前活跃", a.Hours)
	}
}

// DaysDesc 根据天生成描述
func (a *Active) DaysDesc() {
	if a.Days > 30 {
		// 去月生成描述
		a.MonthsDesc()
	} else if a.Days >= 1 && a.Days < 7 {
		a.Description = fmt.Sprintf("%v天前活跃", a.Days)
	} else if a.Days >= 7 && a.Days < 14 {
		a.Description = "一周前活跃"
	} else if a.Days >= 14 && a.Days < 21 {
		a.Description = "两周前活跃"
	} else if a.Days >= 21 && a.Days < 30 {
		a.Description = "三周前活跃"
	}
}

// MonthsDesc 根据月生成描述
func (a *Active) MonthsDesc() {
	if a.Months > 12 {
		// 去年生成描述
		a.YearsDesc()
	} else {
		a.Description = fmt.Sprintf("%v月前活跃", a.Months)
	}
}

// YearsDesc 根据年生成描述
func (a *Active) YearsDesc() {
	if a.Years > 3 {
		a.Description = "很久以前活跃"
	} else {
		a.Description = fmt.Sprintf("%v年前活跃", a.Years)
	}
}
