package handlers

import (
	"context"
	"github.com/boram-gong/apiFlow/common"
	"github.com/boram-gong/apiFlow/operation/db_client"
	"github.com/boram-gong/service/body"
)

func MakeDBClient(ctx context.Context, request interface{}) (interface{}, error) {
	respBody := body.NewCommonResp()
	req := request.(*common.DBClient)
	err := db_client.OperateDB(req)
	if err != nil {
		respBody.FailResp(404, err.Error())
	}
	return respBody, nil
}

func GetDBClient(ctx context.Context, request interface{}) (interface{}, error) {
	if request.(string) == "" {
		respBody := body.NewCommonResp()
		respBody.Data = db_client.GetAllClient()
		return respBody, nil
	} else {
		respBody := body.NewCommonResp()
		respBody.Data = db_client.GetOneClient(request.(string))
		return respBody, nil
	}

}
