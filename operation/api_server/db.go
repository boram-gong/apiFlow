package api_server

import (
	"fmt"
	"github.com/boram-gong/apiFlow/operation"
	dbt "github.com/boram-gong/db_tool"
)

const (
	SqlServerTable = "sql_server"
)

func InsertNewApiServer(key, content string) error {
	s := dbt.InsertSql(
		SqlServerTable,
		[]string{"key", "del", "content"},
		fmt.Sprintf("'%v',0,'%v'", key, content),
	)
	_, err := operation.SelfClient.Exec(s)
	if err != nil {
		return UpdateApiServer(key, content)
	} else {
		return err
	}

}

func UpdateApiServer(key, content string) error {
	s := dbt.UpdateSql(
		SqlServerTable,
		fmt.Sprintf("key='%v'", key),
		[]string{"del=0", "content='" + content + "'"},
	)
	_, err := operation.SelfClient.Exec(s)
	return err
}

func DeleteApiServer(key string) error {
	s := dbt.UpdateSql(
		SqlServerTable,
		fmt.Sprintf("key='%v'", key),
		[]string{"del=1", "content=''"},
	)
	_, err := operation.SelfClient.Exec(s)
	return err
}

func ClearFailureApiServer() {
	_, _ = operation.SelfClient.Exec(dbt.DeleteSql(SqlServerTable, "del!=0"))
}
