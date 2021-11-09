package main

import (
	"fmt"
	"math"
	"strings"
)

//利用set，判断不重复子串
func getMaxChildStr(s string) int{
	if s == ""{
		return 0
	}

	var (
		l, r = 0, 0
		cnt = 1
		mp = make(map[rune]rune)
		newRune = []rune(s)
	)

	for l < len(s)  && r < len(s){
		//存在，移除元素，l++
		if _, ok := mp[newRune[r]]; ok{
			delete(mp, newRune[r])
			l += 1
		}else{
			mp[newRune[r]] = newRune[r]
			r += 1
			cnt = int(math.Max(float64(r - l), float64(cnt)))
		}
	}

	return cnt
}

func getMaxStr(s string) int{
	if s == ""{
		return 0
	}

	var (
		left, right, res = 0, 0, 0
		window = s[left: right]
	)

	for ; right < len(s); right++ {
		if index := strings.IndexByte(window, s[right]); index != -1{
			left = index + 1
		}
		window = s[left: right+1]
		res = int(math.Max(float64(res), float64(len(window))))
	}

	return res
}

func main(){
	fmt.Println(getMaxStr("abbc"))
}
