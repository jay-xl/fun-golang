package basic

import "fmt"

// 声明变量
func DeclaratVariable(){
	// 声明一个变量
	var name string
	// 为声明的变量赋值
	name = "hvkcoder"

	// 声明一个变量的同时为其赋值
	var age = 18

	fmt.Printf("Hello！My name'is %d. I'm %d years old", name, age)
}

// 分组声明
func DeclaratGroupVariable(){
	// 分组声明变量
	var (
		name string
		age int = 18
	)

	name = "hvkcoder"

	fmt.Printf("Hello！My name'is %d. I'm %d years old", name, age)
}

// 简写声明
func SimpleDeclaratVariable(){
	// 简写声明，只能用于局部变量的声明，且必须赋值
	name := "hvkcoder"
	age := 18

	fmt.Printf("Hello！My name'is %d. I'm %d years old", name, age)
}

// 声明多个变量
func DeclaratMultiVariable(){
	// 声明多个变量
	var a , b, c int
	// 为多个变量赋值
	a, b, c = 1 ,2, 3
	fmt.Printf("a = %d , b = %d , c = %d", a, b, c)

	// 声明多个变量并赋值
	var e, f, g int = 4, 5, 6
	fmt.Printf("e = %d , f = %d , g = %d", e, f, g)

	// 简写声明
	h, i, j := 7, 8, 9
	fmt.Printf("h = %d , i = %d , j = %d", h, i, j)
}