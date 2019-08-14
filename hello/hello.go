package main

import (
		"fmt"
	"os"
)

//func main(){
//	var JAVAHOME string
//	JAVAHOME = os.Getenv("GOGCCFLAGS")
//	fmt.Println(JAVAHOME)
//}

func main() {
	//创建一个新的名称为 keys 值为 Hello World! 的环境变量
	err := os.Setenv("GOGCCFLAGS", "-fPIC -m64 -pthread -fmessage-length=0")
	if err != nil {
		//如果出错则输出错误信息
		fmt.Printf("%s", err)
	} else {
		//如果执行成功则输出 set key OK!
		fmt.Println("set key OK!")
		//输出所有环境变量字段和值
		for _, v := range os.Environ() {
			fmt.Println(v)
		}
	}
}