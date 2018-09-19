package basic

import (
	"fmt"
	"unsafe"
)

// 常量定义
func DeclaratConstant() {

	// 显式定义常量
	const _PI float64 = 3.14

	// 隐式定义常量
	const PI = 3.14
}

// 在常量中是用内置表达式
func DeclaratExpression() {
	const (
		_NAME        string = "hvkcoder"
		_NAME_LENGTH int    = len(_NAME)
		_NAME_SIZE          = unsafe.Sizeof(_NAME)
	)
}

func DeclaratGroupConstant() {
	const (
		_NAME string = "hvkcoder"
		_FIRST_NAME
	)

	fmt.Println(_NAME)
	fmt.Println(_FIRST_NAME)
}



// iota 跳值使用法
func JumpValueByIota(){
	const (
		a = 3.14
		b
		c = iota
		d
	)
	fmt.Println(a,b,c,d)
}

// iota 插值使用法
func InsertValueByIota(){
	const (
		a = iota
		b = 3.14
		c = iota
		d
	)
	fmt.Println(a,b,c,d)
}

// iota 表达式隐式使用法
func ExpressionByIota(){
	const (
		B = 1 << (10 * iota)
		KB
		MB
		GB
		TB
		PB
	)

	fmt.Println(B,KB,MB,GB,TB,PB)
}


// iota 单行使用法
func SingleLineByIota(){
	const (
		a , b = iota + 1, iota + 2
		c, d = iota + 3, iota + 4

		// a => 0 + 1
		// b => 0 + 2
		// c => 1 + 3
		// d => 1 + 4
	)
	fmt.Println(a,b,c,d)
}