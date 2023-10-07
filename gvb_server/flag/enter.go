package flag

import (
	sys_flag "flag"
)

type Option struct {
	DB   bool
	User string // -u admin  -u user
	ES   bool   // -es create  -es delete
	Dump string
	Load string
}

/*
解析原理：
flag.Parse() 函数是 flag 包中的一个方法，它用于解析命令行参数并将其写入到注册的标志(flag)中。

当您调用 flag.Parse() 函数时，它会解析当前程序运行时的命令行参数，并将参数与已注册的标志进行匹配。命令行参数通常以空格分隔，并以键值对（如 -flag value）或单独的标志（如 -flag）的形式提供。

例如，假设您的程序名为 myprogram，并且您定义了一个布尔型标志 -verbose，您可以在命令行中输入以下内容来解析命令行参数：
myprogram -verbose
在上面的示例中，-verbose 是一个单独的标志，没有附加值。当您的程序调用 flag.Parse() 函数时，它会将 -verbose 标志的值设置为 true，表示该标志已被指定。


如果您定义的标志需要附加值，您可以在命令行中使用键值对的形式来指定。例如，假设您定义了一个字符串型标志 -name，您可以在命令行中输入以下内容来解析命令行参数：
myprogram -name John
在上面的示例中，-name 是一个键值对形式的标志，John 是其附加的值。当您的程序调用 flag.Parse() 函数时，它会将 -name 标志的值设置为 John，以便您在程序中使用。

通过解析命令行参数并将其与已注册的标志进行匹配，您可以在程序中获取命令行参数的值，并根据需要进行后续处理。
*/

// Parse 解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	// 解析命令行参数写入注册的flag里
	sys_flag.Parse()

	return Option{
		DB: *db,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) (f bool) {
	if option.DB {
		return true
	}
	return false
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
		return
	}
}