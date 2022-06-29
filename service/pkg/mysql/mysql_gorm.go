package mysql

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

const (
	DefaultOrmType    = "gorm"
	DefaultDriverName = "mysql"
)

var (
	DB GormConns
)

//GormConns
type GormConns map[string]*gorm.DB

//Init
func Init(ormType string, db MultipleMysqlConf) (err error) {
	DB, err = NewGORMConns(db)
	if err != nil {
		return err
	}
	return nil
}

//NewGORMConns
func NewGORMConns(confs MultipleMysqlConf) (GormConns, error) {
	conns := make(GormConns)
	for k, v := range confs {
		conn, err := gorm.Open(defaultDriverName, v.Dsn)
		if err != nil {
			return conns, fmt.Errorf(`msg="mysql=%s connection failed." err="%+v"`,
				k,
				err,
			)
		}
		conn.DB().SetConnMaxLifetime(time.Second * 600)
		conn.DB().SetMaxIdleConns(10)
		conn.DB().SetMaxOpenConns(200)

		conn.LogMode(v.Debug)

		conns[k] = conn
	}
	return conns, nil
}
