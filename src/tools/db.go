package tools

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

var x *xorm.Engine



func InitDBConn(tcpConfig *TcpConfig) error {

	dbConnInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", tcpConfig.Db.DbUser, tcpConfig.Db.DbPassword, tcpConfig.Db.DbHost, tcpConfig.Db.DbPort, tcpConfig.Db.UseDbAccountName)
	var err error
	x, err = xorm.NewEngine(tcpConfig.Db.DbDriveName, dbConnInfo)
	if err != nil {
		msg := fmt.Sprintf("Failed to connect to db '%s', err: %s", dbConnInfo, err.Error())
		return errors.New(msg)
	}
	x.SetMaxIdleConns(tcpConfig.Db.DbMaxIdleConn)
	x.SetMaxOpenConns(tcpConfig.Db.DbMaxOpenConn)
	x.ShowSQL(tcpConfig.Db.ShowSql)

	return nil
}

func OrmEngine() *xorm.Engine {
	return x
}
