/**
* @file: utils.go ==> common/utils
* @package: utils
* @author: jingxiu
* @since: 2022/11/8
* @desc: 工具函数
 */

package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	uuid "github.com/satori/go.uuid"
)

func HideString(str string) string {
	if str == "" {
		return "***"
	}
	if len(str) == 2 {
		return string(str[0]) + "*"
	} else {
		return string(str[0]) + "**"
	}
}

//匹配 手机号,邮箱,中文,身份证等等 进行脱敏处理
func HideStar(str string) (result string) {
	if str == "" {
		return "***"
	}
	if strings.Contains(str, "@") {
		// 邮箱
		res := strings.Split(str, "@")
		if len(res[0]) < 3 {
			resString := "***"
			result = resString + "@" + res[1]
		} else {
			res2 := Substr2(str, 0, 3)
			resString := res2 + "***"
			result = resString + "@" + res[1]
		}
		return result
	} else {
		reg := `^1[0-9]\d{9}$`
		rgx := regexp.MustCompile(reg)
		mobileMatch := rgx.MatchString(str)
		if mobileMatch {
			// 手机号
			result = Substr2(str, 0, 3) + "****" + Substr2(str, 7, 11)
		} else {
			nameRune := []rune(str)
			lens := len(nameRune)
			if lens <= 1 {
				result = "***"
			} else if lens == 2 {
				result = string(nameRune[:1]) + "*"
			} else if lens == 3 {
				result = string(nameRune[:1]) + "**"
			} else if lens == 4 {
				result = string(nameRune[:1]) + "***"
			} else if lens > 4 {
				result = string(nameRune[:2]) + "*****"
			}
		}
		return
	}
}
func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	return string(rs[start:end])
}

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

//生成uuid
func UUID() string {
	return uuid.NewV4().String()
}

// RandString 生成随机字符串
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

//随机数字
func RESrand(len int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(rand.Intn(10) + 48)
	}
	return string(bytes)
}

//生成token
func TokenCreate(args ...string) string {
	md5Ctx := md5.New()
	str := ""
	for _, v := range args {
		str += v
	}
	md5Ctx.Write([]byte(fmt.Sprintf("%s%d", str, time.Now().UnixNano())))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// GetLocalIPv4 获取当前服务器IPv4地址
func GetLocalIPv4() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

//对象转json字符串
func JsonStr(v interface{}) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}

//解析JsonStr
func JsonUnmarshal(jsonStr string, v interface{}) {
	_ = json.Unmarshal([]byte(jsonStr), v)
}

// Str2Bytes 字符串转字节切片
func Str2Bytes(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

// Bytes2Str 字节切片转字符串
func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Btoi 布尔值转整形
func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Atoi8 字符串转换成 int8
func Atoi8(s string, d ...int8) int8 {
	i, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}
	return int8(i)
}

// Atoi16 字符串转换成 int8
func Atoi16(s string, d ...int16) int16 {
	i, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}
	return int16(i)
}

// Atoi 字符串转换成 int
func Atoi(s string, d ...int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}

	return i
}

// Atoi32 字符串转换成 int32
func Atoi32(s string, d ...int32) int32 {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}

	return int32(i)
}

// Atoi64 字符串转换成 int64
func Atoi64(s string, d ...int64) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}

	return i
}

// AtoUi8 字符串转换成 uint8
func AtoUi8(s string, d ...uint8) uint8 {
	i, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}
	return uint8(i)
}

// AtoUi16 字符串转换成 uint16
func AtoUi16(s string, d ...uint16) uint16 {
	i, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}
	return uint16(i)
}

// AtoUi 字符串转换成 uint
func AtoUi(s string, d ...uint) uint {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}
	return uint(i)
}

// AtoUi32 字符串转换成 uint32
func AtoUi32(s string, d ...uint32) uint32 {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}
	return uint32(i)
}

// AtoUi64 字符串转换成 uint64
func AtoUi64(s string, d ...uint64) uint64 {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}

	return i
}

// Atof 字符串转换成 float32
func Atof(s string, d ...float32) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}

	return float32(f)
}

// Atof64 字符串转换成 float64
func Atof64(s string, d ...float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		if len(d) > 0 {
			return d[0]
		} else {
			return 0
		}
	}

	return f
}

// I16toA int8 转字符串
func I8toA(i int8) string {
	return strconv.FormatInt(int64(i), 10)
}

// I16toA int16 转字符串
func I16toA(i int16) string {
	return strconv.FormatInt(int64(i), 10)
}

// Itoa int 转字符串
func Itoa(i int) string {
	return strconv.Itoa(i)
}

// I32toA int32 转字符串
func I32toA(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

// I64toA int64 转字符串
func I64toA(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Ui8toA uint8 转字符串
func Ui8toA(i uint8) string {
	return strconv.FormatUint(uint64(i), 10)
}

// Ui16toA uint16 转字符串
func Ui16toA(i uint16) string {
	return strconv.FormatUint(uint64(i), 10)
}

// UitoA uint 转字符串
func UitoA(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}

// Ui32toA uint32 转字符串
func Ui32toA(i uint32) string {
	return strconv.FormatUint(uint64(i), 10)
}

// Ui64toA uint64 转字符串
func Ui64toA(i uint64) string {
	return strconv.FormatUint(i, 10)
}

// F32toA float32 转字符串
func F32toA(f float32) string {
	return F64toA(float64(f))
}

// F64toA float64 转字符串
func F64toA(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func Atoi32Array(in []string) (out []int32) {
	for _, v := range in {
		out = append(out, Atoi32(v))
	}
	return
}

func LitoArray(in string) (out []int32) {
	for _, _vv := range strings.Split(in, "|") {
		if _int := Atoi32(_vv); _int > 0 {
			out = append(out, _int)
		}
	}
	return
}

func LitoUi32Array(in string) (out []uint32) {
	for _, _vv := range strings.Split(in, "|") {
		if _int := AtoUi32(_vv); _int > 0 {
			out = append(out, _int)
		}
	}
	return
}

func LitoUi32ArrayRobot(in string) (out []uint32) {
	for _, _vv := range strings.Split(in, "|") {
		out = append(out, AtoUi32(_vv))
	}
	return
}

func LitoArrayTask(in string) (out []int32) {
	for _, _vv := range strings.Split(in, "|") {
		out = append(out, Atoi32(_vv))
	}
	return
}

func InI32Array(i int32, arr []int32) bool {
	for _, v := range arr {
		if v == i {
			return true
		}
	}
	return false
}

func I32to2f(value int32) (out float64) {
	out, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", float64(value)), 64)
	return
}
func F32to2f(value float32) float32 {
	out, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 32)
	return float32(out)
}
func F64to2f(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
func StringToJsonByte(str string) (b []byte) {
	b, _ = json.Marshal(str)
	return
}

func GetZeroTm() (tm int64) {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}
func DateStrToTimeStamp(str string, tmstr string) (tm int64) {
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s %s", str, tmstr), time.Local)
	tm = theTime.Unix()
	return
}
func TimeStrToTimeStamp(str string) (tm int64) {
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	tm = theTime.Unix()
	return
}

func GetRefreshInterval(refreshType int32) (val int64) {
	t := time.Now()
	refresh_tm := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
	//获取今日刷新时间
	switch refreshType {
	case 1: //每日刷新
		refresh_tm += 3600 * 24
	case 2: //每周刷新
		switch t.Weekday() {
		case time.Sunday:
			refresh_tm += 3600 * 24 * 7
		case time.Monday:
			refresh_tm += 3600 * 24 * 6
		case time.Tuesday:
			refresh_tm += 3600 * 24 * 5
		case time.Wednesday:
			refresh_tm += 3600 * 24 * 4
		case time.Thursday:
			refresh_tm += 3600 * 24 * 3
		case time.Friday:
			refresh_tm += 3600 * 24 * 2
		case time.Saturday:
			refresh_tm += 3600 * 24 * 1
		}
	}
	return refresh_tm - t.Unix()
}

//获取今天日期
func GetToday() (str string) {
	return time.Now().Format("2006-01-02")
}

//时间戳转日期
func TmToDate(tm int64) (str string) {
	return time.Unix(tm, 0).Format("2006-01-02")
}

//

//从数组中删除元素
func DelItemFromIntArr(arr []int, item int) (out []int) {
	length := len(arr)
	for k, v := range arr {
		if v == item {
			if k == length {
				arr = arr[0:k]
			} else {
				arr = append(arr[0:k], arr[k+1:]...)
			}
			break
		}
	}
	return arr
}

func MD5(in string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(in)))
}

//给字符串生成md5
//@params str 需要加密的字符串
//@params salt interface{} 加密的盐
//@return str 返回md5码
func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

//返回一个16位md5加密后的字符串
func Get16MD5Encode(data string, salt ...interface{}) string {
	return Md5Crypt(data, salt)[8:24]
}
