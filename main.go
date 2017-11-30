package main

import (
	_ "sportdatacenter/routers"
)

func foo(data interface{}) (ret interface{}) {
	var jsondata = map[string]interface{}{
		"data":         data,
		"errorMessage": "ok",
	}
	ret = jsondata
	return ret
}

func main() {
	// if beego.BConfig.RunMode == "dev" {
	// 	beego.BConfig.WebConfig.DirectoryIndex = true
	// 	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	// }
	// beego.Run()
	// data := map[string]string{"name": "deng"}
	// fmt.Println(json.Decoder(data))
	// data = make(map[string]interface{})
	// ret := foo(data)
	// ret := foo(data)
	// var age int64
	// fmt.Println(age.(type))
}
