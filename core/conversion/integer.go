package conversion

import "database/sql"

func FormatNullInt16(nullInt16 sql.NullInt16) int {
	if nullInt16.Valid {
		return int(nullInt16.Int16)
	} 
	return -1
}

func FormatNullInt32(nullInt32 sql.NullInt32) int {
	if nullInt32.Valid {
		return int(nullInt32.Int32)
	} 
	return -1
}
