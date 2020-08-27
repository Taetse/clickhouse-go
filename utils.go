package clickhouse

func GetProfileInfo(rows *rows) *ProfileInfo {
	return rows.profileInfo
}
