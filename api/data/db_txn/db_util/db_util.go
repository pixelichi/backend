package db_util

import (
	"database/sql"
	"time"
)

func NullStringToString(str sql.NullString) (toString string) {
	if str.Valid {
		toString = str.String
	} else {
		toString = ""
	}

	return
}

func TimeToString(time time.Time) string {
	return time.Format("2006-01-02T15:04:05Z07:00")
}
