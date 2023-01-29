package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

type Blog struct {
	Path  string
	Title string
	Time  string
}

func (c *IndexController) Get() {
	c.TplName = "index.html"

	blogs := []Blog{
		{"/blog/a", "Hello World", "2017-09-21"},
		{"/blog/b", "Hello World", "2017-07-26"},
		{"/blog/a", "Hello World", "2017-09-21"},
		{"/blog/b", "Hello World", "2017-07-26"},
		{"/blog/a", "Hello World", "2017-09-21"},
		{"/blog/b", "Hello World", "2017-07-26"},
		{"/blog/a", "Hello World", "2017-09-21"},
		{"/blog/b", "Hello World", "2017-07-26"},
	}

	c.Data["Blogs"] = blogs
}
