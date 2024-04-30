package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"botperational/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBConfig struct {
	User            string
	Password        string
	Host            string
	Port            string
	DBName          string
	MaxIdleConn     int
	ConnMaxLifetime int
	MaxOpenConn     int
}

type Sqlx struct {
	*sqlx.DB
	Conf *config.Config `inject:"config"`
}

func (s *Sqlx) Startup() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		s.Conf.Database.User, s.Conf.Database.Password, s.Conf.Database.Host, s.Conf.Database.Port, s.Conf.Database.Name)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	maxIdleConn := 20
	connMaxLifeTime := 1
	maxOpenConn := 100

	if s.Conf.Database.MaxIdleConn != 0 {
		maxIdleConn = s.Conf.Database.MaxIdleConn
	}

	if s.Conf.Database.ConnMaxLifetime != 0 {
		connMaxLifeTime = s.Conf.Database.ConnMaxLifetime
	}

	if s.Conf.Database.MaxOpenConn != 0 {
		maxOpenConn = s.Conf.Database.MaxOpenConn
	}

	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(connMaxLifeTime) * time.Hour)
	db.SetMaxOpenConns(maxOpenConn)

	s.DB = db

	return nil
}

func (s *Sqlx) Shutdown() error {
	sqlDB := s.DB

	return sqlDB.Close()
}

func (s *Sqlx) Ping() error {
	sqlDB := s.DB

	return sqlDB.Ping()
}

func (s *Sqlx) TxOpts() *sql.TxOptions {

	var isolationLevel sql.IsolationLevel
	switch strings.ToLower(s.Conf.Database.TransactionIsolationLevel) {
	case "readuncommitted":
		isolationLevel = sql.LevelReadUncommitted
	case "readcommitted":
		isolationLevel = sql.LevelReadCommitted
	case "writecommitted":
		isolationLevel = sql.LevelWriteCommitted
	case "repeatableread":
		isolationLevel = sql.LevelRepeatableRead
	case "serializable":
		isolationLevel = sql.LevelSerializable
	default:
		isolationLevel = sql.LevelDefault
	}

	return &sql.TxOptions{
		Isolation: isolationLevel,
		ReadOnly:  s.Conf.Database.TransactionReadOnly,
	}
}
