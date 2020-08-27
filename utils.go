package clickhouse

import "database/sql/driver"

func GetProfileInfo(driverRows driver.Rows) *ProfileInfo {
	if localRows, ok := driverRows.(*rows); ok {
		return localRows.profileInfo
	}
	return nil
}
