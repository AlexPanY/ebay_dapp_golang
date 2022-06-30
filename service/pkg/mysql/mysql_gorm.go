package mysql

import (
	"fmt"
	"time"

	m "gorm.io/driver/mysql"
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
		orm, err := gorm.Open(m.Open(v.Dsn), &gorm.Config{})
		if err != nil {
			return conns, fmt.Errorf(`msg="mysql=%s connection failed." err="%+v"`,
				k,
				err,
			)
		}
		conn, err := orm.DB()
		if err != nil {
			return conns, fmt.Errorf(`msg="mysql=%s connection failed." err="%+v"`,
				k,
				err,
			)
		}

		conn.SetConnMaxLifetime(time.Second * 600)
		conn.SetMaxIdleConns(10)
		conn.SetMaxOpenConns(200)

		conns[k] = orm
	}
	return conns, nil
}
