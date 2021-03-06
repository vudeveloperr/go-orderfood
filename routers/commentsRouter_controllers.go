package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-orderfood/controllers:StatisticController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:StatisticController"],
        beego.ControllerComments{
            Method: "GetStatistic",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:StatisticStoreController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:StatisticStoreController"],
        beego.ControllerComments{
            Method: "GetStatisticStore",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-orderfood/controllers:UserController"] = append(beego.GlobalControllerRouter["go-orderfood/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
