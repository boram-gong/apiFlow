package handlers

import (
	"context"
	"github.com/boram-gong/apiFlow/common"
	"github.com/boram-gong/apiFlow/operation/api_server"
	"github.com/boram-gong/apiFlow/operation/db_client"
	dbt "github.com/boram-gong/db_tool"
	"github.com/boram-gong/service/body"
)

func GetApiServer(ctx context.Context, request interface{}) (interface{}, error) {
	respBody := body.NewCommonResp()
	respBody.Data = api_server.GetAllSqlServer()
	return respBody, nil
}

func MakeApiServer(ctx context.Context, request interface{}) (interface{}, error) {
	respBody := body.NewCommonResp()
	req := request.(*common.ServerApiCfg)
	cli, ok := db_client.DBContainer.Load(req.DbClientName)
	if !ok || cli == nil {
		respBody.FailResp(404, "req.DbClientName "+"error")
	} else {
		err := api_server.MakeServer(req, cli.(dbt.DB))
		if err != nil {
			respBody.FailResp(404, err.Error())
		}
	}
	return respBody, nil
}

func ChangeApiServer(ctx context.Context, request interface{}) (interface{}, error) {
	respBody := body.NewCommonResp()
	req := request.(*common.ServerApiCfg)
	cli, ok := db_client.DBContainer.Load(req.DbClientName)
	if !ok || cli == nil {
		respBody.FailResp(404, "req.DbClientName "+"error")
	} else {
		err := api_server.ChangeServerRoute(req, cli.(dbt.DB))
		if err != nil {
			respBody.FailResp(404, err.Error())
		}
	}
	return respBody, nil
}

func DeleteApiServer(ctx context.Context, request interface{}) (interface{}, error) {
	respBody := body.NewCommonResp()
	req := request.(*common.ServerApiPath)
	err := api_server.DeleteServerRoute(req.ServerPort, req.HttpMethod, req.RelativePath)
	if err != nil {
		respBody.FailResp(404, err.Error())
	}

	return respBody, nil
}
