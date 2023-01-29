package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var dbName = "gemini"
var driverName = "mysql"
//user:password@host&port/dbname?charset
var dataSource = "gemini:usergemini@tcp(localhost:3306)/gemini?charset=utf8"
var maxIdle = 30
var maxConn = 30

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql",
		"gemini:usergemini@tcp(localhost:3306)/gemini?charset=utf8")
	orm.RegisterDataBase(
		dbName,
		driverName,
		dataSource,
		maxIdle,
		maxConn)
	orm.RegisterModel(new(Article))
	orm.RunSyncdb("default",true,true)
}

type Article struct {
	Id      int64  `orm:"pk"`
	Title   string `orm:"title"`
	Url     string
	Content string
}

type ArticleManager interface {
	Insert(a *Article) error
	DeleteById(id int64) error
	DeleteByTitle(t string) error
	Update(a *Article) error
	QueryById(id int64) (*Article, error)
	QueryByTitle(t string) (*Article, error)
	QueryAll() ([]*Article, error)
}

type ArticleMySQLManager struct {
}

func NewArticleMySQLManager() *ArticleMySQLManager {
	m := new(ArticleMySQLManager)
	return m
}

func (m *ArticleMySQLManager) Insert(a *Article) error {
	o := orm.NewOrm()
	o.Using(dbName)
	_, err := o.Insert(a)
	if err != nil {
		return err
	}
	return nil
}

func (m *ArticleMySQLManager) DeleteById(id int64) error {
	return nil
}
