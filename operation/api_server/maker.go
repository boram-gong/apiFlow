package api_server

import (
	"errors"
	"github.com/boram-gong/apiFlow/common"
	"github.com/boram-gong/apiFlow/operation"
	"github.com/boram-gong/apiFlow/operation/db_client"
	dbt "github.com/boram-gong/db_tool"
	json_op "github.com/boram-gong/json-decorator/operation"
	"github.com/gin-gonic/gin"
	"sync"
)

var ServerContainer sync.Map

func InitSqlServer() {
	ClearFailureApiServer()
	serverCfg, _ := operation.Query(dbt.SelectFieldsSql(
		SqlServerTable,
		"content",
		"del=0",
	), operation.SelfClient)
	for _, c := range serverCfg {
		var apiCfg common.ServerApiCfg
		if err := common.Decode(c["content"].(string), &apiCfg); err != nil {
			continue
		}
		cli, ok := db_client.DBContainer.Load(apiCfg.DbClientName)
		if ok && cli != nil {
			MakeServer(&apiCfg, cli.(dbt.DB))
		}
	}
}

func GetAllSqlServer() (out []common.ServerApiCfg) {
	serverCfg, _ := operation.Query(dbt.SelectFieldsSql(
		SqlServerTable,
		[]string{"del", "content"},
		"del!=1"), operation.SelfClient)
	for _, c := range serverCfg {
		var apiCfg common.ServerApiCfg
		if err := common.Decode(c["content"].(string), &apiCfg); err != nil {
			continue
		}
		apiCfg.Stat = c["del"]
		out = append(out, apiCfg)
	}
	return
}

// 新增注册路由
func MakeServer(apiCfg *common.ServerApiCfg, client dbt.DB) error {
	var (
		server *ServerStatus
	)
	value, ok := ServerContainer.Load(apiCfg.ServerPort)
	if ok {
		server = value.(*ServerStatus)
	} else {
		server = new(ServerStatus)
		server.Engine = gin.Default()
		server.Lock()
		server.Route = make(map[string]*RouteStatus)
		server.Unlock()
		go func(s *ServerStatus) {
			s.Err = s.Engine.Run(":" + apiCfg.ServerPort)
			if s.Err != nil {
				s.DeleteAllRoute()
			}
		}(server)
	}
	if server.Err != nil {
		return server.Err
	}
	ServerContainer.Store(apiCfg.ServerPort, server)
	err := server.NewRoute(apiCfg, client)
	if err != nil {
		return err
	}

	return InsertNewApiServer(apiCfg.ServerPort+apiCfg.HttpMethod+apiCfg.RelativePath, common.Encode(apiCfg))
}

func ChangeServerRoute(change *common.ServerApiCfg, cli dbt.DB) error {
	server, ok := ServerContainer.Load(change.ServerPort)
	if ok {
		return server.(*ServerStatus).ChangeRoute(change, cli.(dbt.DB))
	} else {
		return errors.New(change.ServerPort + " is not exist")
	}
}

func DeleteServerRoute(port string, httpMethod, relativePath string) error {
	server, ok := ServerContainer.Load(port)
	if ok {
		return server.(*ServerStatus).DeleteRoute(httpMethod, relativePath)
	} else {
		return errors.New(port + " is not exist")
	}
}

func MakeRespBody(querySql string, client dbt.DB, jsonRule string) (interface{}, error) {
	result, err := operation.Query(querySql, client)
	if err != nil {
		return nil, err
	}
	if len(result) == 1 {
		out := result[0]
		if jsonRule != "" {
			err = json_op.DecoratorJsonByRule(jsonRule, out)
		}
		return out, err
	} else {
		out := []map[string]interface{}{}
		for _, v := range result {
			temp := v
			if jsonRule != "" {
				json_op.DecoratorJsonByRule(jsonRule, temp)
			}
			out = append(out, temp)
		}
		return out, nil
	}
}
