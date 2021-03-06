package main

import (
	"net/http"

	"github.com/hprose/hprose-golang/rpc"
	"github.com/jc3wish/Bifrost/hprose_server/serverdo"
)


func main() {
	service := rpc.NewHTTPService()
	service.Debug = true
	service.AddFunction("Insert", serverdo.Insert)
	service.AddFunction("Update", serverdo.Update)
	service.AddFunction("Delete", serverdo.Delete)
	service.AddFunction("ToList", serverdo.ToList)
	service.AddFunction("Check", serverdo.Check)
	http.ListenAndServe(":8881", service)
}
