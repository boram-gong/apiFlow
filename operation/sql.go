package operation

import (
	"fmt"
	"strings"
)

func SelectFieldsSql(table string, fields interface{}, where string) string {
	var (
		sql         string
		selectField string
	)
	switch fields.(type) {
	case string:
		selectField = fields.(string)
	case []string:
		selectField = strings.Join(fields.([]string), ",")

	}
	if where != "" {
		sql = fmt.Sprintf("select %s from %s where %s", selectField, table, where)
	} else {
		sql = fmt.Sprintf("select %s from %s", selectField, table)
	}
	return sql
}

func InsertSql(table string, fields interface{}, values string) string {
	var (
		sql         string
		insertField string
	)
	switch fields.(type) {
	case string:
		insertField = fields.(string)
	case []string:
		insertField = strings.Join(fields.([]string), ",")

	}
	if values != "" {
		sql = fmt.Sprintf("insert into %s (%s) values (%s)", table, insertField, values)
	}
	return sql
}

func UpdateSql(table, where string, change []string) string {
	return fmt.Sprintf(
		"update %s set %v where %s",
		table,
		strings.Join(change, ","),
		where,
	)
}

func DeleteSql(table, where string) string {
	if where != "" {
		return fmt.Sprintf("delete from %s where %s", table, where)
	} else {
		return ""
	}
}
