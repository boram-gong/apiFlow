package server

import (
	"fmt"
	"github.com/boram-gong/apiFlow/common"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/boram-gong/apiFlow/operation/db_client"
	json_comm "github.com/boram-gong/json-decorator/common"
	"github.com/boram-gong/service/svc"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
)

func DecodeNull(c *gin.Context) (interface{}, error) {
	return nil, nil
}

// 解析json转换服务请求体
func DecodeTagJsonReq(c *gin.Context) (interface{}, error) {
	reqBody := &json_comm.JsonReq{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, svc.NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.ConfigFastest.Unmarshal(buf, &reqBody); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, svc.NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}

		switch reqBody.Data.(type) {
		case map[string]interface{}:
			reqBody.JsonMap = reqBody.Data.(map[string]interface{})
		case []interface{}:
			reqBody.JsonSlice = reqBody.Data.([]interface{})
		default:
			return nil, svc.NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
	}
	return reqBody, nil
}

// 解析新增（post）规则请求体
func DecodePostJsonRule(c *gin.Context) (interface{}, error) {
	reqBody := &json_comm.SaveRuleReq{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, svc.NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.ConfigFastest.Unmarshal(buf, &reqBody); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, svc.NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
		if reqBody.Id != 0 {
			return nil, svc.NewError(http.StatusBadRequest, "save id != 0")
		}
	}
	return reqBody, nil
}

// 解析修改（put）规则请求体
func DecodePutJsonRule(c *gin.Context) (interface{}, error) {
	reqBody := &json_comm.SaveRuleReq{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, svc.NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.ConfigFastest.Unmarshal(buf, &reqBody); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, svc.NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
		if reqBody.Id == 0 {
			return nil, svc.NewError(http.StatusBadRequest, "save id == 0")
		}
	}
	return reqBody, nil
}

// 解析参数规则id
func DecodeJsonRuleId(c *gin.Context) (interface{}, error) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		return 0, nil
	} else {
		return id, nil
	}
}

func DecodeDbName(c *gin.Context) (interface{}, error) {
	return c.Query("name"), nil
}

func DecodePostDbClient(c *gin.Context) (interface{}, error) {
	reqBody := &common.DBClient{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, svc.NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.ConfigFastest.Unmarshal(buf, &reqBody); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, svc.NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
		if reqBody.Cfg.MaxOpenConn == 0 {
			reqBody.Cfg.MaxOpenConn = 2
		}
		if reqBody.Cfg.MaxIdleConn == 0 {
			reqBody.Cfg.MaxIdleConn = 8
		}
	}
	return reqBody, nil
}

func DecodePutDbClient(c *gin.Context) (interface{}, error) {
	reqBody := &common.DBClient{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, svc.NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.ConfigFastest.Unmarshal(buf, &reqBody); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, svc.NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
		reqBody.Op = db_client.UPDATE

		if reqBody.Cfg.MaxOpenConn == 0 {
			reqBody.Cfg.MaxOpenConn = 8
		}
		if reqBody.Cfg.MaxIdleConn == 0 {
			reqBody.Cfg.MaxIdleConn = 2
		}
	}
	return reqBody, nil
}

func DecodeDeleteDbClient(c *gin.Context) (interface{}, error) {
	if c.Query("name") == "" {
		return nil, svc.NewError(http.StatusBadRequest, "name is null")
	}
	reqBody := &common.DBClient{
		Name: c.Query("name"),
		Op:   db_client.DELETE,
	}

	return reqBody, nil
}

func DecodeServerApiReq(c *gin.Context) (interface{}, error) {
	reqBody := &common.ServerApiCfg{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, svc.NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.ConfigFastest.Unmarshal(buf, &reqBody); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, svc.NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
		reqBody.HttpMethod = strings.ToTitle(reqBody.HttpMethod)
		reqBody.Sql = strings.ReplaceAll(reqBody.Sql, ";", "")
	}
	return reqBody, nil
}

func DecodeServerApiPathReq(c *gin.Context) (interface{}, error) {
	reqBody := &common.ServerApiPath{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, svc.NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.ConfigFastest.Unmarshal(buf, &reqBody); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, svc.NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
		reqBody.HttpMethod = strings.ToTitle(reqBody.HttpMethod)

	}
	return reqBody, nil
}
