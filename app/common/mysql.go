package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"ldap-server/app/setting"
	"log"
	"runtime"
	"time"
)

// 全局数据库对象
var DB *gorm.DB

func InitDB() {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&%s",
	//	setting.Config.Mysql.Username,
	//	setting.Config.Mysql.Password,
	//	setting.Config.Mysql.Host,
	//	setting.Config.Mysql.Port,
	//	setting.Config.Mysql.Database,
	//	setting.Config.Mysql.Charset,
	//	setting.Config.Mysql.Collation,
	//)
	dsn := "root:123456@tcp(127.0.0.1:3306)/sr1?charset=utf8mb4&parseTime=True&loc=Local"
	//root:12346@tcp(localhost:3306)/ldap?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local&timeout=10000ms
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键(指定外键时不会在mysql创建真实的外键约束)
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy:                           schema.NamingStrategy{SingularTable: true}})
	sqldb, err := db.DB()
	sqldb.SetMaxIdleConns(400)
	sqldb.SetMaxOpenConns(500)
	sqldb.SetConnMaxLifetime(50 * time.Second)
	sqldb.SetConnMaxIdleTime(50 * time.Second)
	go func() {
		for {
			time.Sleep(time.Second * 1)

			// 获取连接池的统计信息
			stats := sqldb.Stats()
			fmt.Println("---------------")
			// 输出连接池的统计信息
			fmt.Printf("MaxOpenConnections: %d\n", stats.MaxOpenConnections)
			fmt.Printf("OpenConnections: %d\n", stats.OpenConnections)
			fmt.Printf("InUse: %d\n", stats.InUse)
			fmt.Printf("Idle: %d\n", stats.Idle)
			numGoroutines := runtime.NumGoroutine()
			fmt.Printf("当前协程数量：%d\n", numGoroutines)
		}
	}()
	if err != nil {
		log.Panicf("初始化mysql数据库异常: %v", err)
	}
	// 开启mysql日志
	if setting.Config.Mysql.LogMode {
		db.Debug()
	}

	DB = db

}
