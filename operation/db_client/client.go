package db_client

import (
	"errors"
	"fmt"
	"github.com/boram-gong/apiFlow/cfg"
	comm "github.com/boram-gong/apiFlow/common"
	"github.com/boram-gong/apiFlow/operation"
	dbt "github.com/boram-gong/db_tool"
	dbt_ms "github.com/boram-gong/db_tool/mysql"
	dbt_pg "github.com/boram-gong/db_tool/pg"
	"github.com/boram-gong/json-decorator/common"
	"sync"
)

const (
	DbClientTable = "db_client"

	DELETE = "delete"
	UPDATE = "update"

	POSTGRE = "postgre"
	MYSQL   = "mysql"
)

var (
	DBContainer sync.Map
)

func GetAllClient() []*cfg.DBClient {
	var respData []*cfg.DBClient
	result, _ := operation.Query(operation.SelectFieldsSql(DbClientTable, "content", ""), operation.SelfClient)
	for _, m := range result {
		var data cfg.DBClient
		if err := comm.Decode(common.Interface2String(m["content"]), &data); err != nil {
			continue
		}
		respData = append(respData, &data)
	}
	return respData
}

func GetOneClient(name string) *cfg.DBClient {
	var data cfg.DBClient
	_, ok := DBContainer.Load(name)
	if ok {
		result, _ := operation.Query(operation.SelectFieldsSql(DbClientTable, "content", "name='"+name+"'"), operation.SelfClient)
		for _, m := range result {
			if err := comm.Decode(common.Interface2String(m["content"]), &data); err != nil {
				continue
			}
		}
		return &data
	} else {
		return nil
	}
}

func InitAllClient() {
	all := GetAllClient()
	for _, cli := range all {
		if cli.State == "success" {
			client, fail := MakeDB(cli.Source, cli.Cfg)
			if fail != nil {
				cli.State = "fail: " + fail.Error()
				operation.SelfClient.Exec(operation.UpdateSql(
					DbClientTable,
					fmt.Sprintf("name='%v'", cli.Name),
					[]string{"content='" + comm.Encode(cli) + "'"},
				))
			} else {
				DBContainer.Store(cli.Name, client)
			}
		} else {
			DBContainer.Store(cli.Name, nil)
		}
	}
}

func NewDB(c *cfg.DBClient) error {
	_, ok := DBContainer.Load(c.Name)
	if ok {
		switch c.Op {
		case DELETE:
			if _, err := operation.SelfClient.Exec(operation.DeleteSql(DbClientTable, fmt.Sprintf("name='%v'", c.Name))); err != nil {
				return err
			} else {
				DBContainer.Delete(c.Name)
				return nil
			}
		case UPDATE:
			DBContainer.Delete(c.Name)
		default:
			return errors.New(c.Name + " operation exist")
		}
	}
	if c.Op == "" || c.Op == UPDATE {
		client, fail := MakeDB(c.Source, c.Cfg)
		if fail != nil {
			if c.Op == UPDATE {
				DBContainer.Store(c.Name, nil)
				c.State = "fail: " + fail.Error()
				if _, err := operation.SelfClient.Exec(operation.UpdateSql(
					DbClientTable,
					fmt.Sprintf("name='%v'", c.Name),
					[]string{"content='" + comm.Encode(c) + "'"},
				)); err != nil {
					return err
				}
			}
			return errors.New("connect operation fail")
		}
		c.State = "success"
		if c.Op == UPDATE {
			if _, err := operation.SelfClient.Exec(operation.UpdateSql(
				DbClientTable,
				fmt.Sprintf("name='%v'", c.Name),
				[]string{"content='" + comm.Encode(c) + "'"},
			)); err != nil {
				return err
			}
		} else {
			if _, err := operation.SelfClient.Exec(
				operation.InsertSql(
					DbClientTable,
					[]string{"name", "content"},
					fmt.Sprintf("'%v','%v'", c.Name, comm.Encode(c)))); err != nil {
				return err
			}
		}
		DBContainer.Store(c.Name, client)
		return nil
	} else {
		return errors.New(c.Name + "is not exist")
	}
}

func MakeDB(source string, cfg *dbt.CfgDB) (dbt.DB, error) {
	switch source {
	case POSTGRE:
		return dbt_pg.NewPgClient(cfg)
	case MYSQL:
		return dbt_ms.NewMysqlClient(cfg)
	default:
		return nil, errors.New(source + " does not support")
	}
}
