package basic

import "fmt"

// 算数运算符
func ArithmeticOperator() {
	count := 0
	fmt.Printf("%d + 1 = %d \n", count, count+1)
	fmt.Printf("%d - 1 = %d \n", count, count-1)
	fmt.Printf("%d * 1 = %d \n", count, count*1)
	fmt.Printf("%d / 1 = %d \n", count, count/1)
	fmt.Println(count, "% 1 = ", count%1)
	count++
	fmt.Println("count 自增值 = ", count)
	count--
	fmt.Println("count 自减值 = ", count)
}

// 关系运算符
func RelationalOperator() {
	fmt.Println("1 == 1", 1 == 1, "1 == 2", 1 == 2)
	fmt.Println("1 != 1", 1 != 1, "1 != 2", 1 != 2)
	fmt.Println("1 > 1", 1 > 1, "1 < 2", 1 < 2)
	fmt.Println("1 >= 1", 1 >= 1, "1 <= 0", 1 <= 0)
}

// 按位运算符
func BitOperator() {
	fmt.Println(" 3 & 4 = ", 3&4)
	fmt.Println(" 3 | 4 = ", 3|4)
	fmt.Println(" 25 ^ 3 = ", 25 ^ 3)
	fmt.Println(" 3 << 3 = ", 3 << 3)
	fmt.Println(" 3 >> 3 = ", 3 >> 3)
}
