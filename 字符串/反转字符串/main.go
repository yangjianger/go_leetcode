package main

import "fmt"

func reverseStr(s string) string{
	if s == ""{
		return s
	}

	//1.去掉字符串中的空格
	var (
		lens = 0
		spec = true
		cur = 0
		strByte = []rune(s)
	)

	//去掉空格并获取有效字符串长度
	for i := 0; i < len(strByte); i++{
		if strByte[i] != ' '{
			strByte[cur] = strByte[i]
			cur++
			spec = false
		}else if spec == false{
			strByte[cur] = ' '
			cur++
			spec = true
		}
	}

	lens = cur
	if spec{
		lens = lens - 1
	}

	if lens <= 0 {
		return ""
	}

	//反转整体字符串
	newStrByte := strByte[0: lens]
	_reverseStr(newStrByte, 0, lens)
	
	//反转部分字符串
	specIndex := -1
	for i := 0; i < len(newStrByte); i++ {
		if newStrByte[i] == ' '{
			_reverseStr(newStrByte,  specIndex+1, i)
			specIndex = i
		}
	}
	_reverseStr(newStrByte,  specIndex+1, lens)

	return string(newStrByte)
}

//这里是左闭右开的 [li, ri)
func _reverseStr(str []rune, li, ri int){
	ri -= 1
	for li < ri{
		str[ri], str[li] = str[li], str[ri]
		li += 1
		ri -= 1
	}
}

func main(){
	fmt.Printf("666_%s_666", reverseStr("   "))
}
