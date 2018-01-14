package app

import (
	"log"
	_ "github.com/lib/pq"
	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	"os"
)

type Book struct {
	Id int `json:"id" orm:"auto"`
	Title string `json:"title" orm:"size(64)"`
	Author string `json:"author" orm:"size(64)"`
}

type ENVConfig struct {
	dbname string
	host string
	user string
	password string
}

var ormObject orm.Ormer

func (u *Book) TableName() string {
	return "book"
}

func init() {
	var force bool
	log.Println(gin.Mode())
	if gin.Mode() == "test" {
		force = true
	} else {
		force = false
	}
	env := getEnvConfig()
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase(
		"default",
		"postgres",
		"user="+env.user+" dbname="+env.dbname+" host="+env.host+" sslmode=disable password="+env.password,
		20)
	orm.RegisterModel(new(Book))
	err := orm.RunSyncdb("default", force, true)
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

func getEnvConfig() ENVConfig {
	var env ENVConfig
	env.dbname = os.Getenv("PUBOOK_DBNAME")
	env.host = os.Getenv("PUBOOK_HOST")
	env.user = os.Getenv("PUBOOK_USER")
	env.password = os.Getenv("PUBOOK_PASSWORD")
	log.Println(env)
	return env
}
