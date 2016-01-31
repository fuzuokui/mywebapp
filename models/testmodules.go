package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"     //引入beego的orm
	_ "github.com/go-sql-driver/mysql" //引入beego的mysql驱动
)

type Bar struct {
	Id      int
	BarId   int
	BarName string `orm:"size(60)"`
	BarDesc string `orm:"size(100)"`
}

type Article struct {
	Id            int
	ArticleId     int `orm:"pk;auto"`
	BarId         int
	MediaDigestId int
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@/bookbar?charset=utf8")
	orm.SetMaxIdleConns("default", 30)           //设置数据库最大空闲连接
	orm.SetMaxOpenConns("default", 30)           //设置数据库最大连接数
	orm.RegisterModelWithPrefix("go_", new(Bar)) //注册模型并使用表前缀
	orm.RegisterModelWithPrefix("go_", new(Article))
}

//自动建表
func createTable() {
	name := "default"                          //数据库别名
	force := false                             //不强制建数据库
	verbose := true                            //打印建表过程
	err := orm.RunSyncdb(name, force, verbose) //建表
	if err != nil {
		fmt.Println(err)
	}
}

func (this *Bar) Add(barId int, barName string, barDesc string) (int64, error) {
	o := orm.NewOrm()
	o.Using("bookbar") //使用数据库，默认default
	createTable()      //开启自动建表
	bar := Bar{BarId: barId, BarName: barName, BarDesc: barDesc}
	id, err := o.Insert(&bar)
	return id, err
}

func (this *Bar) QueryAllBar() ([]Bar, error) {
	o := orm.NewOrm()
	var data []Bar
	_, err := o.QueryTable("go_bar").Filter("id__gt", 8).All(&data)
	return data, err
}
