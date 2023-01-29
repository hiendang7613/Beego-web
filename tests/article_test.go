package test

import (
	"testing"
	"Gemini/models"
	"math/rand"
	"github.com/astaxie/beego/logs"
	"runtime"
	"path/filepath"
	"fmt"
	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	fmt.Println("path:" + filepath.Join(file, ".."+string(filepath.Separator)))
	fmt.Println("apppath:" + apppath)
	beego.TestBeegoInit(apppath + "/../")
}

func TestInsert(t *testing.T) {
	m := models.NewArticleMySQLManager()
	a := new(models.Article)
	a.Id = rand.Int63()
	a.Title = "Hello World"
	a.Url = "/blog/hello-word"
	a.Content = "This is a test post"
	err := m.Insert(a)
	if err != nil {
		logs.Error("Err insert article: ", err)
	}
}
