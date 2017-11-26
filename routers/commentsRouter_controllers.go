package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["sportdatacenter/controllers:ObjectController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:ObjectController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:ObjectController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:ObjectController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:ObjectController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:UserController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:UserController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:UserController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:UserController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:UserController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:UserController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:UserController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:WalkdataController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:WalkdataController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:WalkdataController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:WalkdataController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:WalkdataController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:WalkdataController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:WalkdataController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:WalkdataController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:WalkdataController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:WalkdataController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["sportdatacenter/controllers:WalkdataController"] = append(beego.GlobalControllerRouter["sportdatacenter/controllers:WalkdataController"],
		beego.ControllerComments{
			Method: "GetBatchData",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
