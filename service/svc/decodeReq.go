package svc

import (
	"fmt"
	"github.com/boram-gong/apiFlow/cfg"
	"github.com/boram-gong/apiFlow/operation/db_client"
	"github.com/boram-gong/json-decorator/common/body"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func DecodeNull(c *gin.Context) (interface{}, error) {
	return nil, nil
}

// 解析json转换服务请求体
func DecodeTagJsonReq(c *gin.Context) (interface{}, error) {
	reqBody := &body.JsonReq{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.ConfigFastest.Unmarshal(buf, &reqBody); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}

		switch reqBody.Data.(type) {
		case map[string]interface{}:
			reqBody.JsonMap = reqBody.Data.(map[string]interface{})
		case []interface{}:
			reqBody.JsonSlice = reqBody.Data.([]interface{})
		default:
			return nil, NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
	}
	return reqBody, nil
}

// 解析新增（post）规则请求体
func DecodePostJsonRule(c *gin.Context) (interface{}, error) {
	reqBody := &body.SaveRuleReq{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.ConfigFastest.Unmarshal(buf, &reqBody); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
		if reqBody.Id != 0 {
			return nil, NewError(http.StatusBadRequest, "save id != 0")
		}
	}
	return reqBody, nil
}

// 解析修改（put）规则请求体
func DecodePutJsonRule(c *gin.Context) (interface{}, error) {
	reqBody := &body.SaveRuleReq{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.ConfigFastest.Unmarshal(buf, &reqBody); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
		if reqBody.Id == 0 {
			return nil, NewError(http.StatusBadRequest, "save id == 0")
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
	reqBody := &cfg.DBClient{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.ConfigFastest.Unmarshal(buf, &reqBody); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, NewError(http.StatusBadRequest,
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
	reqBody := &cfg.DBClient{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.ConfigFastest.Unmarshal(buf, &reqBody); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, NewError(http.StatusBadRequest,
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
		return nil, NewError(http.StatusBadRequest, "name is null")
	}
	reqBody := &cfg.DBClient{
		Name: c.Query("name"),
		Op:   db_client.DELETE,
	}

	return reqBody, nil
}

func DecodeServerApiReq(c *gin.Context) (interface{}, error) {
	reqBody := &cfg.ServerApiCfg{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.ConfigFastest.Unmarshal(buf, &reqBody); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
		reqBody.HttpMethod = strings.ToTitle(reqBody.HttpMethod)
		reqBody.Sql = strings.ReplaceAll(reqBody.Sql, ";", "")
	}
	return reqBody, nil
}

func DecodeServerApiPathReq(c *gin.Context) (interface{}, error) {
	reqBody := &cfg.ServerApiPath{}
	buf, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = json.ConfigFastest.Unmarshal(buf, &reqBody); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
		reqBody.HttpMethod = strings.ToTitle(reqBody.HttpMethod)

	}
	return reqBody, nil
}
