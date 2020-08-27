package clickhouse

import (
	"database/sql"
	"reflect"
)

func GetRowsBeforeLimit(rows *sql.Rows) uint64 {
	rowsiPtr := reflect.ValueOf(rows).Elem().FieldByName("rowsi")
	rowsi := reflect.Indirect(rowsiPtr).Elem()
	profileInfoPtr := reflect.Indirect(rowsi).FieldByName("profileInfo")
	rowsBeforeLimit := reflect.Indirect(profileInfoPtr).FieldByName("rowsBeforeLimit").Uint()
	return rowsBeforeLimit
}
