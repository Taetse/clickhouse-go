package clickhouse

import (
	"database/sql"
	"reflect"
)

func GetProfileInfo(rows sql.Rows) *ProfileInfo {
	rowsiPtr := reflect.ValueOf(rows).Elem().FieldByName("rowsi")
	rowsi := reflect.Indirect(rowsiPtr).Elem()
	profileInfoPtr := reflect.Indirect(rowsi).FieldByName("ProfileInfo")
	return profileInfoPtr.Interface().(*ProfileInfo)
}
