package main

import (
	_ "base/env"
	//"base/log"
	"base/rpc"
	//_ "github.com/astaxie/beego"
	"log"
)

//var logger = log.GetLogger()

func main() {

	log.SetFlags(log.Ldate|log.Llongfile)


	rpc.UDP()

}
