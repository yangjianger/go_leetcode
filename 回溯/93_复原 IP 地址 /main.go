package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。

有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。

例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。

 

示例 1：

输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
示例 2：

输入：s = "0000"
输出：["0.0.0.0"]
示例 3：

输入：s = "1111"
输出：["1.1.1.1"]
示例 4：

输入：s = "010010"
输出：["0.10.0.10","0.100.1.0"]
示例 5：

输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
 */

func main(){
	fmt.Println(restoreIpAddresses("010010"))
}

func restoreIpAddresses(s string) []string {

	res := make([]string,  0)
	for i := 1; i < 4; i++{
		for j := 1; j < 4; j++{
			for k := 1; k < 4; k++{
				for m := 1; m < 4; m++{
					if i + j + k + m != len(s){
						continue
					}

					a := s[0:i]
					b := s[i:j+i]
					c := s[i+j:i+j+k]
					d := s[i+j+k:]

					if !checkStr(a) || !checkStr(b) || !checkStr(c) || !checkStr(d){
						continue
					}

					if atoi(a) <= 255  && atoi(b) <= 255 && atoi(c) <= 255 && atoi(d) <= 255{
						temp := a + "." + b + "." + c + "." + d
						if len(temp)  == len(s) + 3{
							res = append(res, temp)
						}
					}
				}
			}
		}
	}

	return res
}

func checkStr(s string) bool{
	if strings.Index(s, "0") == 0 && len(s) > 1{
		return false
	}

	return true
}

func atoi(i string) int{
	res, _ := strconv.Atoi(i)
	return res
}