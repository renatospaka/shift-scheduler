package conversion

import "database/sql"

func FormatNullBool(nullBool sql.NullBool) bool {
	if nullBool.Valid {
		return nullBool.Bool
	} 
	return false
}
