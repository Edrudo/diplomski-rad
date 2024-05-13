package mysql

import "github.com/volatiletech/sqlboiler/v4/boil"

// Database is an abstraction of database.
// It can be called to execute database operations.
type Database interface {
	boil.ContextBeginner
	boil.ContextExecutor
}
