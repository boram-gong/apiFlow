package api_server

import (
	"fmt"
	"github.com/boram-gong/apiFlow/operation"
)

const (
	SqlServerTable = "api_server"
)

func InsertNewApiServer(key, content string) error {
	s := operation.InsertSql(
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
	s := operation.UpdateSql(
		SqlServerTable,
		fmt.Sprintf("key='%v'", key),
		[]string{"del=0", "content='" + content + "'"},
	)
	_, err := operation.SelfClient.Exec(s)
	return err
}

func DeleteApiServer(key string) error {
	s := operation.UpdateSql(
		SqlServerTable,
		fmt.Sprintf("key='%v'", key),
		[]string{"del=1", "content=''"},
	)
	_, err := operation.SelfClient.Exec(s)
	return err
}

func ClearFailureApiServer() {
	_, _ = operation.SelfClient.Exec(operation.DeleteSql(SqlServerTable, "del!=0"))
}
