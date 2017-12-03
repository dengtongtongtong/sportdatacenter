package main

import (
	_ "sportdatacenter/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()

	// query := webutils.Query{Method: "GET", Url: "https://httpbin.org/get", Data: map[string]string{"name": "dengtongtong"}, Headers: map[string]string{"agent": "fuck"}}
	// feedback := webutils.CaptureSingle(query, 100*time.Second)
	// fmt.Println(feedback.Error)
	// fmt.Println(feedback.Message)

	// type ColorGroup struct {
	// 	ID     int
	// 	Name   string
	// 	Colors []string
	// }
	// group := ColorGroup{
	// 	ID:     1,
	// 	Name:   "Reds",
	// 	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	// }
	// b, err := json.Marshal(group)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// os.Stdout.Write(b)
}
