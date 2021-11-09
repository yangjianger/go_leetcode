package main

import (
	"fmt"
)

//题目：将n个盘子从A移动到C, 大盘只能在小盘下面
//解题思路：这个其实就是三步
/**
	1.将n-1的盘子移动到B
	2.将第n个盘子移动到C
	3.将n-1盘子移动到C
 */
func main(){
	hanoi(3, "a", "b", "c")
}

func hanoi(n int, p1, p2, p3 string){
	//设置递归基
	if n == 1{
		move(n, p1, p3)
		return
	}
	//1.将n-1的盘子A移动到B
	hanoi(n-1, p1, p3, p2)
	//2.将第n个盘子移动到C
	move(n, p1, p3)
	//3.将n-1盘子从p2移动到C
	hanoi(n-1, p2, p1, p3)

}

func move(no int, p1, p2 string){
	fmt.Printf("将%d号盘子从%s移动到%s\n", no, p1, p2)
}
