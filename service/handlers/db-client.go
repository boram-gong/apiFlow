package handlers

import (
	"context"
	"github.com/boram-gong/apiFlow/cfg"
	"github.com/boram-gong/apiFlow/common"
	"github.com/boram-gong/apiFlow/operation/db_client"
)

func MakeDBClient(ctx context.Context, request interface{}) (interface{}, error) {
	respBody := common.NewCommonResp()
	req := request.(*cfg.DBClient)
	err := db_client.NewDB(req)
	if err != nil {
		respBody.FailResp(404, err.Error())
	}
	return respBody, nil
}

func GetDBClient(ctx context.Context, request interface{}) (interface{}, error) {
	if request.(string) == "" {
		respBody := common.NewCommonResp()
		respBody.Data = db_client.GetAllClient()
		return respBody, nil
	} else {
		respBody := common.NewCommonResp()
		respBody.Data = db_client.GetOneClient(request.(string))
		return respBody, nil
	}

}
