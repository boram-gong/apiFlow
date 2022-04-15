package operation

import (
	dbt "github.com/boram-gong/db_tool"
	dbt_pg "github.com/boram-gong/db_tool/pg"
	"github.com/boram-gong/json-decorator/common"
	json "github.com/json-iterator/go"
	"log"
	"strings"
)

var (
	SelfClient dbt.DB
	pCfg       = &dbt.CfgDB{
		Host:        "114.67.78.94",
		Port:        5432,
		User:        "postgres",
		Password:    "Wayz2022",
		Database:    "response_adapter",
		MaxIdleConn: 5,
		MaxOpenConn: 20,
	}
)

func InitDB() {
	var err error
	SelfClient, err = dbt_pg.NewPgClient(pCfg)
	if err != nil {
		log.Fatalln(err)
	}
}

func Query(querySql string, db dbt.DB) (result []map[string]interface{}, err error) {
	rows, err := db.QueryX(querySql)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		m := map[string]interface{}{}
		if e := rows.MapScan(m); e != nil {
			continue
		}
		for k, v := range m {
			if strings.Contains(common.Interface2String(v), "}") {
				temp := make(map[string]interface{})
				if json.UnmarshalFromString(common.Interface2String(v), &temp) == nil {
					m[k] = temp
				}
			}
		}
		result = append(result, m)

	}
	return result, nil
}
