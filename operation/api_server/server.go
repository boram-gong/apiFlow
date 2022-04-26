package api_server

import (
	"errors"
	"fmt"
	"github.com/boram-gong/apiFlow/common"
	"github.com/boram-gong/apiFlow/operation"
	dbt "github.com/boram-gong/db_tool"
	"github.com/boram-gong/service/body"
	"github.com/gin-gonic/gin"
	"strings"
	"sync"
)

const (
	UrlArg = "url"
)

type ServerStatus struct {
	sync.RWMutex
	Engine *gin.Engine
	Route  map[string]*RouteStatus
	Err    error
}

func (s *ServerStatus) NewRoute(apiCfg *common.ServerApiCfg, client dbt.DB) error {
	s.Lock()
	defer s.Unlock()
	route, ok := s.Route[apiCfg.HttpMethod+apiCfg.RelativePath]
	if ok {
		if !route.Valid {
			route.Valid = true
			route.Db = client
			route.Cfg = apiCfg
		}
		return errors.New(apiCfg.HttpMethod + " " + apiCfg.RelativePath + " exist")
	} else {
		route = new(RouteStatus)
		route.Db = client
		route.Cfg = apiCfg
		route.Valid = true
		s.Engine.Handle(apiCfg.HttpMethod, apiCfg.RelativePath, func(c *gin.Context) {
			if !route.Valid {
				c.JSON(404, nil)
			} else {
				var (
					body     = body.NewCommonResp()
					querySql = route.Cfg.Sql
					e        error
					result   interface{}
				)
				if len(route.Cfg.Args) != 0 {
					for _, arg := range route.Cfg.Args {
						if arg.Source == UrlArg {
							if c.Query(arg.Key) != "" {
								querySql = strings.ReplaceAll(querySql, "@"+arg.Key, c.Query(arg.Key))
							} else {
								body.FailResp(401, "arg '"+arg.Key+"' is null")
								break
							}
						}
					}
				}
				querySql, e = ArgDeal(querySql, route.Cfg.Args, c)
				if e != nil {
					body.FailResp(400, e.Error())
				}
				if body.Code == 200 {
					result, e = MakeRespBody(querySql, route.Db, apiCfg.JsonRule)
					if e != nil {
						body.FailResp(400, e.Error())
					} else {
						body.Data = result
					}
				}
				c.JSON(200, body)
			}
		})
		s.Route[apiCfg.HttpMethod+apiCfg.RelativePath] = route
	}
	return nil
}

func (s *ServerStatus) ChangeRoute(change *common.ServerApiCfg, cli dbt.DB) error {
	s.Lock()
	defer s.Unlock()
	route, ok := s.Route[change.HttpMethod+change.RelativePath]
	if ok {
		if route.Cfg.DbClientName != change.DbClientName && change.DbClientName != "" {
			if cli != nil {
				route.Cfg.DbClientName = change.DbClientName
				route.Db = cli
			}
		}
		if change.Sql != "" && route.Cfg.Sql != change.Sql {
			route.Cfg.Sql = change.Sql
		}
		if route.Cfg.JsonRule != change.JsonRule {
			route.Cfg.JsonRule = change.JsonRule
		}
		route.Cfg.Args = change.Args
		return UpdateApiServer(route.Cfg.ServerPort+change.HttpMethod+change.RelativePath, common.Encode(route.Cfg))
	} else {
		return errors.New(change.HttpMethod + " " + change.RelativePath + " is not exist")
	}
}

func (s *ServerStatus) DeleteRoute(httpMethod, relativePath string) error {
	s.Lock()
	defer s.Unlock()
	route, ok := s.Route[httpMethod+relativePath]
	if ok {
		if route.Valid == false {
			return errors.New(httpMethod + " " + relativePath + " is not valid")
		}
		_ = DeleteApiServer(route.Cfg.ServerPort + httpMethod + relativePath)
		route.Cfg = nil
		route.Valid = false
		return nil
	} else {
		return errors.New(httpMethod + " " + relativePath + " is not exist")
	}
}

func (s *ServerStatus) DeleteAllRoute() {
	s.Lock()
	defer s.Unlock()
	for _, r := range s.Route {
		r.Valid = false
		delSql := dbt.UpdateSql(
			SqlServerTable,
			fmt.Sprintf("key='%v'", r.Cfg.ServerPort+r.Cfg.HttpMethod+r.Cfg.RelativePath),
			[]string{"del=2"},
		)
		operation.SelfClient.Exec(delSql)
	}
}

type RouteStatus struct {
	Cfg   *common.ServerApiCfg
	Db    dbt.DB
	Valid bool
}

func ArgDeal(s string, args []common.ApiArg, c *gin.Context) (string, error) {
	for _, arg := range args {
		if arg.Source == UrlArg {
			if c.Query(arg.Key) != "" {
				s = strings.ReplaceAll(s, "@"+arg.Key, c.Query(arg.Key))
			} else {
				return s, errors.New("arg '" + arg.Key + "' is null")
			}
		}
	}
	return s, nil
}
