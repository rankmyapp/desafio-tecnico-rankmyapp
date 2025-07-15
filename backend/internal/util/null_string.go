package util

import "database/sql"

func ToNullString(val string) sql.NullString {
	return sql.NullString{String: val, Valid: val != ""}
}
