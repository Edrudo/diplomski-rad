package bootstrap

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"http3-server-poc/cmd/api/config"
)

const (
	mysqlConnectionFormat = "%s:%s@tcp(%s)/%s"
)

func newMysqlConnection(logger *zap.Logger) *sql.DB {
	var sqlDb *sql.DB
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			mysqlConnectionFormat,
			config.Cfg.MysqlConfig.Username,
			config.Cfg.MysqlConfig.Password,
			config.Cfg.MysqlConfig.DatabaseAddress,
			config.Cfg.MysqlConfig.DatabaseName,
		),
	)
	if err != nil {
		logger.Warn(errors.Wrap(err, "cannot connect to mysql DB").Error())
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	sqlDb = db
	logger.Info("connected to mysql DB successfully")

	return sqlDb
}
