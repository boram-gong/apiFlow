package main

import (
	"github.com/boram-gong/apiFlow/operation"
	"github.com/boram-gong/apiFlow/operation/api_server"
	"github.com/boram-gong/apiFlow/operation/db_client"
	"github.com/boram-gong/apiFlow/server"
	json_rule "github.com/boram-gong/json-decorator/rule"
	"time"
)

func main() {
	operation.InitDB()
	json_rule.ReAllRule(operation.SelfClient)
	go func() {
		for {
			time.Sleep(10 * time.Minute)
			json_rule.ReAllRule(operation.SelfClient)
		}
	}()
	db_client.InitAllClient()
	api_server.InitSqlServer()
	server.Run("29999", "")
}
