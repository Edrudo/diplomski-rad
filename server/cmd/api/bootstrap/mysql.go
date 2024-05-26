package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/cenkalti/backoff"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"http3-server-poc/cmd/api/config"
)

const (
	mysqlConnectionFormat = "%s:%s@tcp(localhost:%s)/%s"
)

func newMysqlConnection(logger *zap.Logger) *sql.DB {
	var sqlDb *sql.DB
	err := backoff.Retry(
		func() error {

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
				return err
			}

			err = db.Ping()
			if err != nil {
				logger.Warn(errors.Wrap(err, "mysql DB ping failed").Error())
				return err
			}

			sqlDb = db
			logger.Info("connected to mysql DB successfully")
			return nil
		},
		backoff.NewExponentialBackOff(),
	)
	if err != nil {
		panic(err)
	}

	return sqlDb

}
