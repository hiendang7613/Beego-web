package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"strings"
)

var sessionName string = "admin"

type AdminController struct {
	beego.Controller
}

type Sess struct {
	Username string
	Password string
}

func (c *AdminController) Get() {
	if !c.hasLogin() {
		c.Redirect("/login", http.StatusFound)
		return
	}
	c.TplName = "admin.tpl"
	v := c.GetSession(sessionName).(string)
	sesses := strings.Split(v, " ")
	username := sesses[0]
	password := sesses[1]
	c.Data["Session"] = &Sess{Username: username, Password: password}
}

func (c *AdminController) hasLogin() bool {
	v := c.GetSession(sessionName)
	return v != nil
}
