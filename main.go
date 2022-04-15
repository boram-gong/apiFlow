package main

import (
	"github.com/boram-gong/apiFlow/operation"
	"github.com/boram-gong/apiFlow/operation/api_server"
	"github.com/boram-gong/apiFlow/operation/db_client"
	"github.com/boram-gong/apiFlow/operation/json_rule"
	"github.com/boram-gong/apiFlow/service/svc/server"
	"time"
)

func main() {
	operation.InitDB()
	json_rule.ReAllRule()
	go func() {
		for {
			time.Sleep(10 * time.Minute)
			json_rule.ReAllRule()
		}
	}()
	db_client.InitAllClient()
	api_server.InitSqlServer()
	server.Run()
}
