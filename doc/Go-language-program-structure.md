## Go 语言的项目结构

&ensp;&ensp; 一般,一个 Go 项目在 GOPATH 下，会有如下三个目录

<pre>
.        
├── bin   // 存放编译后的可执行文件
├── pkg   // 存放编译后的包文件
└── src   // 存放项目源文件
</pre>

一般情况下，bin 和 pkg 目录可以不创建，go 命令会自动构建（如 go install），只需要创建 src 目录即可。


## Go 程序的一般结构

```go
// 当前程序的包名
package main

// 导入其他包
import "fmt"

// 常量的声明
const _PI float64 = 3.14

// 全局变量声明
var title string = "Go 语言学习笔记"

// 一般类型声明
type newType int

// 结构声明
type Student struct{
	
}

// 接口的声明
type ILearn interface{
	
}

// 由 main 作为程序入口的启动点
func main(){
	fmt.Println("Hello Golang")
}

```

- Go 程序是通过 ```package``` 来组织的（与 python 类似），**```package``` 是 最基本的分发单位 和 工程管理中依赖关系的体现；**
- 每个 Go 语言源代码文件开头必须拥有一个 ```pakcage``` 声明，表示源码文件所属代码包。默认情况下，除 ```main``` 的 ```package``` 包外，其他的包名对应文件夹名称；
- 要生成 Go 语言的可执行程序，必须有 ```main``` 的 ```package``` 包，且必须在该包下面有 ```main()``` 函数；
- **同一个路径下只能存在一个 ```package```，一个 ```package``` 可以拆分成多个源文件**；
- 通过 ```import``` 关键字来导入其他非 ```main``` 包；
- 通过 ```const``` 关键字来定义常量；
- 通过在函数体外部使用 ```var``` 关键字定义全局变量；
- 通过 ```type``` 关键字来进行结构（```struct```）或接口（```interface```）的声明；
- 通过 ```func``` 关键字来声明函数；


## Go 语言中 import 详解

- ```import``` 语句可以导入源代码文件中所依赖的 ```package``` 包，导入包后可以使用 ```<PackageName>.<FuncName>``` 对包中的函数进行调用；
- 如果导入包之后未调用其中的函数或类型将会报编译错误；
- ```import``` 可以使用以下两种方式：
 
  - 单行导入
  
  ```go
    package main
    
    import "fmt"
    import "os"
    import "time"
    import "io"
  ```
  
  - 多行导入
  
  ```go
  package main

  import (
    "fmt"
    "os"
    "time"
    "io"
  )
  
  ```
 
- 如果一个 ```main``` 包导入其他包，包将被顺序导入；
- 如果导入的包依赖其他包（如：包B）,会首先导入包B，然后初始化包B中的常量和变量，最后如果包B中有```init``` 函数，将会自动执行 ```init``` 函数；
- 所有包导入完成后才会对 ```main``` 中变量和常量进行初始化，然后执行 ```main``` 的 ```init``` 函数（如果存在），最后才会执行 ```main``` 函数；
- 如果一个包被导入多次，则该包只会被导入一次；


在使用 ```import``` 导入 ```package``` 包时，可以为其设置别名：

- 自定义别名：

```go
package main

import io "fmt"

func main(){
	io.Println("Hello Golang")
}
```

```go
package main

import (
	io "fmt"
)

func main(){
	io.Println("Hello Golang")
}
```
- 点（```.```）标识的导入包后，调用该包中的函数时，可以省略报名前缀名称（不建议使用）

```go
package main

import . "fmt"

func main(){
	Println("Hello Golang")
}
```

- 下划线（```_```）标识符导入包时，并不是导入整个包，而是执行该包中的 ```init``` 函数，因此无法通过包名来调用包中的其他函或属性。**使用下划线（```_```）操作通常是未来注册包里的引擎，外部可以方便的使用**；

**以上三点不可以同时使用。**