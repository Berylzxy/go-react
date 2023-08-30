package mysql

import (
	"fmt"
	"go-react/model"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitMysql() *gorm.DB {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	//获取ini的数据
	username := cfg.Section("mysql").Key("username").String()
	password := cfg.Section("mysql").Key("password").String()
	ip := cfg.Section("mysql").Key("ip").String()
	port := cfg.Section("mysql").Key("port").String()
	database := cfg.Section("mysql").Key("database").String()
	charset := cfg.Section("mysql").Key("charset").String()
	parseTime := cfg.Section("mysql").Key("parseTime").String()
	loc := cfg.Section("mysql").Key("loc").String()
	//user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local

	//改ini的数据
	//cfg.Section("").Key("app_name").SetValue("testt")
	//cfg.SaveTo("./conf/app.ini")
	dsn := username + ":" + password +
		"@tcp(" + ip + ":" + port + ")/" + database +
		"?charset=" + charset + "&parseTime=" + parseTime +
		"&loc=" + loc
	fmt.Println("-------------", dsn)
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&model.User{}, &model.UserDetail{})
	u1 := model.User{
		Id:       1,
		Userid:   1,
		UserName: "admin",
		PassWord: "123456",
	}
	DB.Create(&u1)
	return DB
}
