package handlers

import (
	"context"
	"github.com/boram-gong/apiFlow/operation/json_rule"
	json_body "github.com/boram-gong/json-decorator/common/body"
	json_op "github.com/boram-gong/json-decorator/operation"
	"github.com/boram-gong/service/body"
)

// json转化器服务
func JsonDecorator(ctx context.Context, request interface{}) (interface{}, error) {
	respBody := body.NewCommonResp()
	reqBody := request.(*json_body.JsonReq)
	if reqBody.JsonMap != nil {
		respJson := reqBody.JsonMap
		err := json_op.DecoratorJsonByRule(reqBody.Name, respJson)
		if err != nil {
			respBody.FailResp(400, err.Error())
		} else {
			respBody.Data = respJson
		}
	} else if reqBody.JsonSlice != nil {
		var respList []interface{}
		for _, j := range reqBody.JsonSlice {
			if err := json_op.DecoratorJsonByRule(reqBody.Name, j); err != nil {
				respBody.FailResp(400, err.Error())
				return respBody, nil
			}
			respList = append(respList, j)
		}
		respBody.Data = respList
	} else {
		respBody.FailResp(404, "no data")
	}
	return respBody, nil
}

// 读取json转换规则服务
func ReadJsonRule(ctx context.Context, request interface{}) (interface{}, error) {
	if request.(int) == 0 {
		respBody := body.NewCommonResp()
		respBody.Data = json_rule.GetAllRule()

		return respBody, nil
	} else {
		oneRule := json_rule.GetOneRule(request.(int))
		respBody := body.NewCommonResp()
		respBody.Data = oneRule
		return respBody, nil
	}
}

// 存储json转换规则服务
func SaveRule(ctx context.Context, request interface{}) (interface{}, error) {
	respBody := body.NewCommonResp()
	saveData := request.(*json_body.SaveRuleReq)
	if err := json_rule.SaveRule(saveData); err != nil {
		respBody.FailResp(500, err.Error())
	}
	respBody.Data = json_rule.ReAllRule()
	return respBody, nil
}

// 删除json转换规则服务
func DeleteRule(ctx context.Context, request interface{}) (interface{}, error) {
	respBody := body.NewCommonResp()
	if err := json_rule.DeleteRule(request.(int)); err != nil {
		respBody.FailResp(400, err.Error())
	}
	respBody.Data = json_rule.ReAllRule()
	return respBody, nil
}

// 重置json转换规则服务
func ReRule(ctx context.Context, request interface{}) (interface{}, error) {
	json_rule.ReAllRule()
	respBody := body.NewCommonResp()
	return respBody, nil
}
