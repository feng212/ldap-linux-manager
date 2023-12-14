package main

import (
	"fmt"
	"ldap-server/app/common"
	"ldap-server/app/model"
	"ldap-server/app/router"
	"ldap-server/app/setting"
	"net/http"
	"time"
)

func main() {
	setting.InitConfig()
	common.InitDB()
	go gorun()
	server := &http.Server{
		Addr:         "localhost:" + "8082",
		Handler:      router.InitRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}

func gorun() {
	// 获取当前协程数量
	for j := 0; j < 20; j++ {
		time.Sleep(time.Second * 1)

		for i := 0; i <= 152; i++ {
			go search()
		}

	}
	fmt.Println("退出............................")
	fmt.Println("退出............................")
	fmt.Println("退出............................")
	fmt.Println("退出............................")
	fmt.Println("退出............................")
	fmt.Println("退出............................")
	fmt.Println("退出............................")
	fmt.Println("退出............................")
	fmt.Println("退出............................")
	fmt.Println("退出............................")
	fmt.Println("退出............................")
	fmt.Println("退出............................")
	fmt.Println("退出............................")
}

func search() {
	list := make([]model.Test, 0)

	common.DB.Table("test").Find(&list)

}
