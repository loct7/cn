package cn

func testConnection() {
	var dbWork DatabaseWork
	_ = dbWork.DB.Ping()
}
