package json_rule

import (
	"errors"
	"fmt"
	"github.com/boram-gong/apiFlow/operation"
	dbt "github.com/boram-gong/db_tool"
	"github.com/boram-gong/json-decorator/common/body"
	"github.com/boram-gong/json-decorator/rule"
	json "github.com/json-iterator/go"
)

const (
	JsonRuleTable = "rule"
)

func ReAllRule() []body.SaveRuleReq {
	ruleMap := rule.NewAllRuleSafeMap()
	var respData []body.SaveRuleReq
	result, _ := operation.Query(dbt.SelectFieldsSql(JsonRuleTable, "*", ""), operation.SelfClient)
	for _, m := range result {
		var data []body.UserRule
		if err := json.UnmarshalFromString(dbt.Interface2String(m["rule"]), &data); err != nil {
			continue
		}
		respData = append(respData, body.SaveRuleReq{
			Id:        dbt.Interface2Int(m["id"]),
			Name:      dbt.Interface2String(m["rule_name"]),
			Rules:     data,
			Stat:      dbt.Interface2Int(m["stat"]),
			StartTime: dbt.Interface2String(m["start_time"]),
			EndTime:   dbt.Interface2String(m["end_time"]),
		})
		for _, d := range data {
			r := &rule.Rule{
				Key:       d.Key,
				Operation: d.Operation,
				Content:   d.Content,
				Stat:      dbt.Interface2Int(m["stat"]),
				StartTime: dbt.Interface2String(m["start_time"]),
				EndTime:   dbt.Interface2String(m["end_time"]),
			}
			ruleName := fmt.Sprintf("%v", m["rule_name"])
			ruleMap.UnSafeStore(ruleName, r)
		}
	}
	rule.AllRule.Store(ruleMap)
	return respData
}

func GetAllRule() []body.SaveRuleReq {
	var respData []body.SaveRuleReq
	result, _ := operation.Query(dbt.SelectFieldsSql(JsonRuleTable, "*", ""), operation.SelfClient)
	for _, m := range result {
		var data []body.UserRule
		if err := json.UnmarshalFromString(dbt.Interface2String(m["rule"]), &data); err != nil {
			continue
		}
		respData = append(respData, body.SaveRuleReq{
			Id:        dbt.Interface2Int(m["id"]),
			Name:      dbt.Interface2String(m["rule_name"]),
			Rules:     data,
			Stat:      dbt.Interface2Int(m["stat"]),
			StartTime: dbt.Interface2String(m["start_time"]),
			EndTime:   dbt.Interface2String(m["end_time"]),
		})
	}
	return respData
}

func GetOneRule(id int) body.SaveRuleReq {
	var respData body.SaveRuleReq
	result, _ := operation.Query(dbt.SelectFieldsSql(JsonRuleTable, "*", fmt.Sprintf("id=%v", id)), operation.SelfClient)
	for _, m := range result {
		var data []body.UserRule
		if err := json.UnmarshalFromString(dbt.Interface2String(m["rule"]), &data); err != nil {
			continue
		}
		respData = body.SaveRuleReq{
			Id:        dbt.Interface2Int(m["id"]),
			Name:      dbt.Interface2String(m["rule_name"]),
			Rules:     data,
			Stat:      dbt.Interface2Int(m["stat"]),
			StartTime: dbt.Interface2String(m["start_time"]),
			EndTime:   dbt.Interface2String(m["end_time"]),
		}
		break
	}
	return respData
}
func SaveRule(data *body.SaveRuleReq) error {
	saveSql := ""
	rules, err := json.Marshal(data.Rules)
	if err != nil {
		return err
	}
	if data.Id != 0 {
		change := []string{
			fmt.Sprintf("rule_name='%v'", data.Name),
			fmt.Sprintf("rule='%s'", string(rules)),
			fmt.Sprintf("stat=%v", data.Stat),
			fmt.Sprintf("start_time='%v'", data.StartTime),
			fmt.Sprintf("end_time='%v'", data.EndTime),
		}
		saveSql = dbt.UpdateSql(JsonRuleTable, fmt.Sprintf("id=%v", data.Id), change)
	} else {
		fields := []string{"rule_name", "rule", "stat", "start_time", "end_time"}
		values := fmt.Sprintf("'%v','%v',%v,'%v','%v'",
			data.Name,
			string(rules),
			data.Stat,
			data.StartTime,
			data.EndTime,
		)
		saveSql = dbt.InsertSql(JsonRuleTable, fields, values)
	}
	_, err = operation.SelfClient.Exec(saveSql)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRule(id int) error {
	if id == 0 {
		return errors.New("id is null")
	}
	_, err := operation.SelfClient.Exec(dbt.DeleteSql(JsonRuleTable, fmt.Sprintf("id=%v", id)))
	if err != nil {
		return err
	}
	return nil
}
