package main


func isValid(s string) bool {

	var (
		stack = []string{}
		tmpS = ""
	)

	for i := 0; i < len(s); i++{
		if string(s[i]) == "(" || string(s[i]) == "[" || string(s[i]) == "{"{
			stack = append(stack, string(s[i]))
		}else{
			//出栈
			tmpS = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if string(s[i]) == ")" && tmpS != "("{
				return false
			}

			if string(s[i]) == "]" && tmpS != "["{
				return false
			}

			if string(s[i]) == "}" && tmpS != "{"{
				return false
			}
		}
	}

	return len(stack) == 0
}

func main(){

}


