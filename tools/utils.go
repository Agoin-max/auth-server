package tools

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

const DATE_TEMPLATE = "2006-01-02 15:04:05"

// 时间戳转换为日期
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format(DATE_TEMPLATE)
}

// 日期转换为时间戳 2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	t, err := time.ParseInLocation(DATE_TEMPLATE, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// 获取时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// 获取当前日期
func GetDate() string {
	return time.Now().Format(DATE_TEMPLATE)
}

// 获取年月日
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

// MD5加密
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}
