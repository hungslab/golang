## install

`brew install go`

## command

```text
go build hello

在src目录或hello目录下执行 go build hello，只在对应当前目录下生成文件。

go install hello

在src目录或hello目录下执行 go install hello，会把编译好的结果移动到 $GOPATH/bin。

go run hello

在src目录或hello目录下执行 go run hello，不生成任何文件只运行程序。

go fmt hello

在src目录或hello目录下执行 go run hello，格式化代码，将代码修改成标准格式。
```

## 目录结构
bin  
存放编译后可执行的文件。

pkg  
存放编译后的应用包。

src  
存放应用源代码。

例如：
```text
├─ code  -- 代码根目录
│  ├─ bin
│  ├─ pkg
│  ├─ src
│     ├── hello
│         ├── hello.go
```

## Usage

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

命令行执行：`go run demo.go` 

## 常量声明

常量，在程序编译阶段就确定下来的值，而程序在运行时无法改变该值。

- 单个常量声明
  - 第一种：const 变量名称 数据类型 = 变量值  
    如果不赋值，使用的是该数据类型的默认值。
  - 第二种：const 变量名称 = 变量值  
    根据变量值，自行判断数据类型。

- 多个常量声明

  - 第一种：const 变量名称,变量名称 ... ,数据类型 = 变量值,变量值 ...
  - 第二种：const 变量名称,变量名称 ... = 变量值,变量值 ...

## 变量声明

- 单个变量声明

  - 第一种：var 变量名称 数据类型 = 变量值  
    如果不赋值，使用的是该数据类型的默认值。
  - 第二种：var 变量名称 = 变量值  
    根据变量值，自行判断数据类型。
  - 第三种：变量名称 := 变量值  
    省略了 var 和数据类型，变量名称一定要是未声明过的。

- 多个变量声明
  - 第一种：var 变量名称,变量名称 ... ,数据类型 = 变量值,变量值 ...
  - 第二种：var 变量名称,变量名称 ... = 变量值,变量值 ...
  - 第三种：变量名称,变量名称 ... := 变量值,变量值 ...

## 输出方法

1. fmt.Print：输出到控制台（仅只是输出）
2. fmt.Println：输出到控制台并换行
3. fmt.Printf：仅输出格式化的字符串和字符串变量（整型和整型变量不可以）
4. fmt.Sprintf：格式化并返回一个字符串，不输出。












