package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
	_ "modernc.org/sqlite"
)

var (
	Ins      *sqlx.DB
	rootPath = "/root/.dnat"
)

func initWorkspace() error {
	if err := os.MkdirAll(rootPath, 0700); err != nil {
		return err
	}
	return nil
}

func InitDB() error {
	driverName := "sqlite"
	dsn := rootPath + "/" + "database.db?_pragma=busy_timeout(50000)&_pragma=journal_mode(WAL)"
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return err
	}

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if err := initWorkspace(); err != nil {
		return err
	}

	file, err := os.OpenFile(rootPath+"/"+"dnat.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return err
	}
	zlogger := zerolog.New(file).With().Timestamp().Logger()

	loggerOptions := []sqldblogger.Option{
		sqldblogger.WithSQLQueryFieldname("sql"),
		sqldblogger.WithWrapResult(false),
		sqldblogger.WithExecerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithQueryerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug),
	}

	db = sqldblogger.OpenDriver(dsn, db.Driver(), zerologadapter.New(zlogger), loggerOptions...)

	Ins = sqlx.NewDb(db, "sqlite")
	if err := Ins.Ping(); err != nil {
		return err
	}
	if _, err := Ins.Exec(schema); err != nil {
		return err
	}
	return nil
}

func Tx(f func(tx *sqlx.Tx) error) (err error) {
	tx := Ins.MustBegin()
	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
				log.Printf("rollback failed: %s\n", err)
			}
			panic(p)
		} else if err != nil {
			if err := tx.Rollback(); err != nil {
				log.Printf("rollback failed: %s\n", err)
			}
		} else {
			if err := tx.Commit(); err != nil {
				log.Printf("commit failed: %s\n", err)
			}
		}
	}()

	err = f(tx)
	return err
}
