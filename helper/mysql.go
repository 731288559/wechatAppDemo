package helper

import (
	"database/sql"
	"demo/conf"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func initDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		conf.RConf.Mysql.Username,
		conf.RConf.Mysql.Password,
		conf.RConf.Mysql.Network,
		conf.RConf.Mysql.Addr,
		conf.RConf.Mysql.Port,
		conf.RConf.Mysql.Database,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
		// fmt.Printf("Open mysql failed,err:%v\n", err)
		// return
	}

	var sqlDB *sql.DB
	sqlDB, err = DB.DB()
	sqlDB.SetConnMaxLifetime(time.Duration(conf.RConf.Mysql.ConnMaxLifetime) * time.Second) //最大连接周期，超过时间的连接就close
	sqlDB.SetMaxOpenConns(conf.RConf.Mysql.MaxOpenConns)                                    //设置最大连接数
	sqlDB.SetMaxIdleConns(conf.RConf.Mysql.MaxIdleConns)                                    //设置闲置连接数
}
