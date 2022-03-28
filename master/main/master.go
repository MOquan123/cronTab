package main

import "runtime"

func InitEv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var(
		err error
	)

	//初始化线程
	InitEv()
	//启动HTTP API服务
	if err = master.InitApiServer();err != nil{
		
	}
}