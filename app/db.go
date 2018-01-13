package app

import (
	"log"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego/orm"
)

type Book struct {
	Id int `json:"id" orm:"auto"`
	Title string `json:"title" orm:"size(64)"`
	Author string `json:"author" orm:"size(64)"`
}

var ormObject orm.Ormer

func (u *Book) TableName() string {
	return "book"
}

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase(
		"default",
		"postgres",
		"user=postgres dbname=pubook-db host=127.0.0.1 sslmode=disable",
		20)
	orm.RegisterModel(new(Book))
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatal(err)
	}
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func GetOrmObject() orm.Ormer {
	return orm.NewOrm()
}