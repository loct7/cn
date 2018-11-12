package cn

import "database/sql"

type DatabaseWork struct {
	DB  *sql.DB
	Cfg *Configuration
}

// Connect to DB

func (dbWork *DatabaseWork) Init(cfg *Configuration) {
	dbWork.Cfg = cfg
	var err error
	dbWork.DB, err = sql.Open("mysql", dbWork.Cfg.User+":"+dbWork.Cfg.Password+"@tcp("+dbWork.Cfg.Host+":"+dbWork.Cfg.Port+")/"+dbWork.Cfg.Database)
	dbWork.DB.SetMaxOpenConns(20)
	dbWork.DB.SetMaxIdleConns(20)
	CheckError(err)

}
