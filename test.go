package cn

import (
	"fmt"
)

//Description:
//
func testConnection() {
	var dbWork DatabaseWork
	_ = dbWork.DB.Ping()
	fmt.Println("Connected successfully!")
}
