package main

import "strconv"

func restoreIpAddresses(s string) (ans []string) {
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}
	var backtrace func(int, string, []string)
	backtrace = func(index int, ip string, current []string) {
		if index == len(s) {
			if len(current) == 4 {
				c := generateIp(current)
				// 多了三个dot
				if len(c) == len(s)+3 {
					ans = append(ans, c)
				}
			}
			return
		}
		// 加当前字符
		ip += string(s[index])
		index++
		// 唯一0
		if len(ip) == 1 && ip == "0" {
			backtrace(index, "", append(current, ip))
			return
		}
		// 此时已不合理 剪枝
		if len(ip) >= 4 {
			return
		}
		// 不会越界
		ipN, _ := strconv.Atoi(ip)
		if ipN > 255 {
			return
		}
		if ipN <= 255 {
			backtrace(index, "", append(current, ip))
		}
		backtrace(index, ip, current)
	}
	backtrace(0, "", []string{})
	return
}
func generateIp(ipArr []string) (ip string) {
	for i := 0; i < len(ipArr); i++ {
		ip += ipArr[i]
		ip += "."
	}
	// 去掉最后的
	ip = ip[:len(ip)-1]
	return
}
