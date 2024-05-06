package shorter

import (
	"github.com/twmb/murmur3"
	"math/rand"
	"strings"
)

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func V2S(val uint32) (url string) {
	url = ""
	for val > 0 {
		url = string(charset[val%62]) + url
		val /= 62
	}
	return
}

func Shorten(url string) string {
	hasher := murmur3.New32()
	hasher.Write([]byte(url))
	val := hasher.Sum32()
	return V2S(val)
}

func Add1(length int) string {
	var sb strings.Builder
	sb.Grow(length)
	// 随机数种子
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}

	return sb.String()
}
