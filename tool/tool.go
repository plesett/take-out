package tool

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"regexp"
)

/**
 * 验证手机号
 */
func IsMobile(mobile string) bool  {
	result, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, mobile)
	if result {
		return true
	} else {
		return false
	}
}

/**
 * 生成32位 md5 字符串
 */
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(h.Sum(nil)))
}

/**
 * 生成 guid 字符串
 */
func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b);err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}