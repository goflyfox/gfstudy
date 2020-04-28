package hello

import (
	"gfcli/app/model/user"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

// Hello is a demonstration route handler for output "Hello World!".
func Hello(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
	entity, err := user.FindOne("login_name = ?", "admin")
	if err != nil {
		glog.Error(err)
		r.Response.Writeln("err")
		r.Exit()
	}
	r.Response.Writeln(entity.Id)
	r.Response.Writeln(entity)
}
