## Go 语言简介

&ensp;&ensp;Go 是一门开源、支持并发、垃圾回收的编译型系统编程语言，从 2007 年末由 Robert Griesemer，Rob Pike，Ken Thompson 主持开发，后来还加入了 lan Lance Taylor，Russ Cox 等人，并最终在 2009 年 11 月 开源，在 2012 年早些时候发布了 Go 1 稳定版本。


## Go 语言的主要特点

- 类型安全 和 内存安全；
- 以非常直观和极低代价的方案实现高并发；
- 高效的垃圾回收机制（内置 runtime）；
- 快速编译（同时解决 C 语言中头文件太多的问题）；
- 为多核计算机提供性能提升的方案；
- UTF-8 编码支持；

## Go 语言的应用

- 服务器编程：处理日志、数据打包、文件系统等；
- 分布式系统：数据库处理器，中间件等；
- 网络编程：目前使用最多最广泛的一块，Web 应用、API 应用等；
- 云平台：目前云平台逐步采用 Go 实现；


## Go 语言中的常用命令

- go get：获取远程包（需要提前安装 git 或 hg）；
- go run：直接运行程序；
- go build：项目发布；
- go fmt：格式化源码（部分 IDE 在保存时自动调用）；
- go install：编译包文件并编译整个程序；
- go test：运行测试文件；
- go doc：查看文档；
- go help：查看 go 命令行；


## Go 语言中的关键字，标识符

&ensp;&ensp;Go 语言中保留关键字只有 25 个

| break | default | func | interface | select |
| --- | --- | --- | --- | --- |
| case | defer | go | map | struct |
| chan | else | goto | package | switch |
| const | fallthrough | if | range | type |
| continue | for | import | return | var |


&ensp;&ensp;Go 语言中有 36 个预定的标识符，其中包括基础数据类型和系统内嵌函数
  
| append | bool | byte | cap | close | complex |
| --- | --- | --- | --- | --- | --- |
| complex64 | complex128 | copy | false | float32 | float64 |
| copy | int | int8 | int16 | int32 | int64 |
| imag | uint | uint8 | uint16 | uint32 | uint64 | 
| uintprt | iota | len | new | nil | panic | 
| recover | print | println | real | string | TRUE | 


## Go 语言中的注释方法

- //：单行注释

- /* */：多行注释

## Go 语言中可见性规则

&ensp;&ensp; Go 语言中，使用大小写来决定该常量、变量、类型、接口、结构是否可以被外部所调用：**根据约定，函数名首字母 小写 即为 ```private```，函数名首字母 大写 即为 ```plublic```**。