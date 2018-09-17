## Go 语言的变量与常量 

### 变量声明，初始化与赋值 （[示例代码](https://github.com/SilenceHVK/fun-golang/blob/master/src/basic/variable.go)）

- 变量的声明格式：```var <变量名称> [变量类型]```；
- 变量的赋值格式：```<变量名称> = <值，表达式，函数等>```；
- 声明和赋值同时进行：```var <变量名称> [变量类型] = <值，表达式，函数等>```；
- 分组声明的格式

```go
package basic

import "fmt"

func DeclaratGroupVariable(){
	// 分组声明变量
	var (
		name string
		age int = 18
	)

	name = "hvkcoder"

	fmt.Printf("Hello！My name'is %d. I'm %d years old", name, age)
}
```

- 简写声明格式：```<变量名称> := <值，表达式，函数等>```，只能在局部变量声明时使用，并且必须赋值；
- 声明多个变量，需要以 ```,``` 隔开；

```go
package basic

import "fmt"

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
```

**声明全局变量时，不能使用简写声明，且必须使用 ```var``` 关键字。变量名是 下划线（```_```) 表示忽略。**

### 常量定义的形式，类型范围（[示例代码](https://github.com/SilenceHVK/fun-golang/blob/master/src/basic/constant.go)）
