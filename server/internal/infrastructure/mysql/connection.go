package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/pkg/errors"
)

// Connection is a struct containing a connection information to the database.
type Connection struct {
	DB *sql.DB
}

// DsnConnection is a struct containing data on how to connect to the database
type DsnConnection struct {
	User     string
	Password string
	Host     string
	DbName   string
	UseSSL   string
}

func (dsn *DsnConnection) logString() string {
	return fmt.Sprintf(
		"user=%s password=X host=%s dbname=%s sslmode=%s",
		dsn.User, dsn.Host, dsn.DbName, dsn.UseSSL,
	)
}

// NewConnection opens a new connection to the database by using the given DSN information.
// It creates and returns a new instance of Connection.
func NewConnection(
	dsn DsnConnection,
	maxIdleConnections int,
	maxOpenConnections int,
	connectionMaxLifetime time.Duration,
	retryInterval time.Duration,
) (*Connection, error) {
	const maxRetries = 10

	errctx := func(err error) error {
		return errors.WithMessagef(err, "auto migrate with DSN: %v", dsn.logString())
	}

	// WARNING: Take care not to log DSN connection string/structure since
	// it contains credentials for the database
	dsnConnection := fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s sslmode=%s",
		dsn.User, dsn.Password, dsn.Host, dsn.DbName, dsn.UseSSL,
	)

	var db *sql.DB

	err := backoff.Retry(
		func() error {
			var err error
			db, err = sql.Open("postgres", dsnConnection)
			if err != nil {
				return errors.Wrap(err, "connecting to the DB")
			}

			err = db.Ping()
			return errors.Wrap(err, "pinging the DB")
		},
		backoff.WithMaxRetries(
			backoff.NewConstantBackOff(retryInterval),
			maxRetries,
		),
	)

	if err != nil {
		return nil, errctx(err)
	}

	db.SetMaxIdleConns(maxIdleConnections)
	db.SetMaxOpenConns(maxOpenConnections)
	db.SetConnMaxLifetime(connectionMaxLifetime)

	return &Connection{
		DB: db,
	}, nil
}
