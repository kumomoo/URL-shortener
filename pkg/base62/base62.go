package base62

import (
	"math"
	"strings"
)

// 62进制转换

// 为了避免被人恶意请求，可以打乱下面的字符串
// const base62Str = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
// const base62Str = `J0rs1205TUV8IW7D9aBdXecCfghiMQj3klmop6qtuvbcwx4zAEFGHKLNnPRYSZy`

var (
	base62Str string
)

// MustInit 要使用base62这个包必须要调用该函数完成初始化
func MustInit(bs string) {
	if len(bs) == 0 {
		panic("need base string!")
	}
	base62Str = bs
}

// 10进制数转为62进制字符串
func Int2String(seq uint64) string {
	if seq == 0 {
		return string(base62Str[0])
	}

	bl := []byte{}
	for seq > 0 {
		mod := seq % 62
		div := seq / 62
		bl = append(bl, base62Str[mod])
		seq = div
	}

	// 最后把倒的数反转一下
	return string(reverse(bl))
}

// 62进制字符串转为10进制数
func String2Int(s string) (seq uint64) {
	bl := []byte(s)
	bl = reverse(bl)
	// 从右往左遍历
	for idx, b := range bl {
		base := math.Pow(62, float64(idx))
		seq += uint64(strings.Index(base62Str, string(b))) * uint64(base)
	}
	return seq
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
