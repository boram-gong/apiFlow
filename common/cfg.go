package common

import dbt "github.com/boram-gong/db_tool"

type ServerApiCfg struct {
	ServerApiPath
	SqlApiCfg
	JsonRule string      `json:"json_rule"`
	Args     []ApiArg    `json:"url_args"`
	Stat     interface{} `json:"stat,omitempty"`
}

type ServerApiPath struct {
	ServerPort   string `json:"port"`
	HttpMethod   string `json:"method"`
	RelativePath string `json:"path"`
}

type SqlApiCfg struct {
	DbClientName string `json:"db_name"`
	Sql          string `json:"sql"`
}

type ApiArg struct {
	Source    string `json:"source"`
	Key       string `json:"key"`
	Operation string `json:"op"`
}

type DBClient struct {
	Source string     `json:"source"`
	Name   string     `json:"name"`
	State  string     `json:"state"`
	Op     string     `json:"-"`
	Cfg    *dbt.CfgDB `json:"cfg"`
}
