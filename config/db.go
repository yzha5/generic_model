package config

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
)

type DbConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	Database  string
	IsRunning bool
}

var (
	DbMasterList = []DbConfig{
		{
			Host:      "127.0.0.1",
			Port:      3306,
			User:      "root",
			Pwd:       "8905130",
			Database:  "gen_model",
			IsRunning: true,
		},
	}

	DbMaster = DbMasterList[0]

	cfg = mysql.Config{
		User:                 DbMaster.User,
		Passwd:               DbMaster.Pwd,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", DbMaster.Host, DbMaster.Port),
		DBName:               DbMaster.Database,
		Collation:            "utf8mb4_general_ci",
		Loc:                  Local,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	MysqlDsn = cfg.FormatDSN()
)
