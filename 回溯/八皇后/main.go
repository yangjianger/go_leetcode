package main

import (
	"fmt"
	"math"
	"strings"
)

var cols []int
var ways = 0

var  (
	cols2 []bool
	leftTop []bool
	rightTop []bool
)

/**
解题思路： 主要思想是回溯+剪枝
 */
func main(){
	n := 4
	cols = make([]int, n)
	cols2 = make([]bool, n)
	leftTop = make([]bool,  n << 2 - 1)
	rightTop = make([]bool, len(leftTop))

	place2(0)

	fmt.Println(ways)
}

//第一种方式好了解一些
func place2(row int){
	if row == len(cols2){
		ways += 1
		show()
		return
	}

	var ltIndex = 0
	var rtIndex = 0
	for col := 0; col < len(cols2); col++ {
		if cols2[col]{
			continue
		}
		ltIndex = row - col + len(cols2) - 1
		if leftTop[ltIndex]{
			continue
		}
		rtIndex = row + col
		if rightTop[rtIndex]{
			continue
		}

		cols2[col] = true
		leftTop[ltIndex] = true
		rightTop[rtIndex] = true
		cols[row] = col
		place2(row + 1)

		//回溯，说明走不通， 或者这种方法已经走完,重置状态
		cols2[col] = false
		leftTop[ltIndex] = false
		rightTop[rtIndex] = false

	}
}

func place(row int){

	if row == len(cols){
		ways += 1
		show()
		return
	}

	for col := 0; col < len(cols); col++ {
		//剪枝， 判断列上，对角上有没有值，没有就设置值
		if isValid(row, col){
			cols[row] = col

			//这里结束之后就用到了回溯的思想，比如：执行完之后，就会从下一列进行判断
			place(row+1)
		}
	}
}

func isValid(row,  col int) bool{
	for i := 0; i < row; i++ {
		//如果同一行相等，返回
		if cols[i] == col{
			return false
		}
		//判断斜率， x1-x2 / y1-y2 == 1
		if float64(row - i) == math.Abs(float64(col - cols[i])){
			return false
		}
	}
	return true
}

func show(){
	item := 0
	for i := 0; i < len(cols); i++ {
		for j := 0; j < len(cols); j++ {
			item = 0
			if j == cols[i]{
				item = 1
			}
			fmt.Printf("%d\t", item)
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("--", 10))
}
