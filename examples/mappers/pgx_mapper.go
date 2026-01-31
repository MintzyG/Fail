package mappers

import (
	"errors"

	"fail"

	"github.com/jackc/pgconn"
)

var (
	SQLUniqueViolation = fail.ID(0, "SQL", 0, true, "SQLUniqueViolation") // 0_SQL_0000_S
	SQLForeignKey      = fail.ID(0, "SQL", 1, true, "SQLForeignKey")      // 0_SQL_0001_S
	SQLUnknownError    = fail.ID(0, "SQL", 2, false, "SQLUnknownError")   // 0_SQL_0002_D
)

var (
	ErrSQLUniqueViolation = fail.Form(SQLUniqueViolation, "unique violation", false, nil)
	ErrSQLForeignKey      = fail.Form(SQLForeignKey, "foreign key violation", false, nil)
	ErrSQLUnknownError    = fail.Form(SQLUnknownError, "unknown error", false, nil)
)

type PGXMapper struct{}

func (m *PGXMapper) Name() string  { return "pgx" }
func (m *PGXMapper) Priority() int { return 100 }

// Map Generic → Generic OR Fail
func (m *PGXMapper) Map(err error) (error, bool) {
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		return nil, false
	}

	switch pgErr.Code {
	case "23505": // unique_violation
		return ErrSQLUniqueViolation, true
	case "23503": // foreign_key_violation
		return ErrSQLForeignKey, true
	default:
		return ErrSQLUnknownError, true
	}
}

// MapFromFail fail.Error → Generic
func (m *PGXMapper) MapFromFail(fe *fail.Error) (error, bool) {
	// Convert to something infrastructure understands
	return errors.New(fe.Message), true
}

// MapToFail Generic → fail.Error
func (m *PGXMapper) MapToFail(err error) (*fail.Error, bool) {
	return ErrSQLUnknownError, true
}
