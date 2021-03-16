package main

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if validIPv4(IP) {
		return "IPv4"
	}
	if validIPv6(IP) {
		return "IPv6"
	}
	return "Neither"
}
func validIPv4(IP string) bool {
	ipArr := strings.Split(IP, ".")
	if len(ipArr) != 4 {
		return false
	}
	for _, v := range ipArr {
		// 长度判断 第一位不能是0
		if len(v) < 1 || len(v) > 3 || (v[0] == '0' && len(v) != 1) {
			return false
		}
		// 转换为数字 范围0-255
		num, error := strconv.Atoi(v)
		if error != nil || num < 0 || num > 255 {
			return false
		}
	}
	return true
}
func validIPv6(IP string) bool {
	// 按冒号分割
	ipArr := strings.Split(IP, ":")
	if len(ipArr) != 8 {
		return false
	}
	for _, v := range ipArr {
		// 长度判断 第一位不能是0
		if len(v) < 1 || len(v) > 4 {
			return false
		}
		for i := 0; i < len(v); i++ {
			if !isDigitOrLetter(v[i]) {
				return false
			}
		}
	}
	return true
}

func isDigitOrLetter(a byte) bool {
	return (a >= '0' && a <= '9') || (a >= 'a' && a <= 'f') || (a >= 'A' && a <= 'F')
}
