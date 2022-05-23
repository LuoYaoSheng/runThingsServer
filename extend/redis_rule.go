package extend

import (
	"encoding/json"
	"github.com/LuoYaoSheng/runThingsConfig/model"
	service "github.com/LuoYaoSheng/runThingsServer/core"
	"log"
)

// RuleFromRedis 获取设备告警规则
func RuleFromRedis(sn, code string) []model.Rule {
	log.SetFlags(log.Llongfile)

	var objRules []model.Rule
	var key string
	// 获取 sn 对应规则
	key = sn + "_rule"
	snValue, _ := service.GetRdValue(key)
	var snRules []model.Rule
	if len(snValue) > 0 {
		err := json.Unmarshal([]byte(snValue), &snRules)
		if err != nil {
			log.Println(err)
			return objRules
		}
	}
	// 获取 code 对应规则
	key = code + "_rule"
	codeValue, _ := service.GetRdValue(key)
	var codeRules []model.Rule
	if len(codeValue) > 0 {
		err := json.Unmarshal([]byte(codeValue), &codeRules)
		if err != nil {
			log.Println(err)
			return objRules
		}
	}

	rules_ := append(snRules, codeRules...) // 一定要 snRules在前，重复时好保留
	objRules = RemoveRepByLoop(rules_)
	return objRules
}

// RemoveRepByLoop 通过两重循环过滤重复元素
func RemoveRepByLoop(slc []model.Rule) []model.Rule {
	var result []model.Rule // 存放结果
	for i := range slc {
		flag := true
		for j := range result {

			log.Println(slc[i].Content)
			var slcMap []model.RuleContent
			err := json.Unmarshal([]byte(slc[i].Content), &slcMap)
			if err != nil {
				log.Println(err)
				break
			}

			var resMap []model.RuleContent
			err = json.Unmarshal([]byte(result[j].Content), &resMap)
			if err != nil {
				log.Println(err)
				break
			}

			if len(slcMap) == len(resMap) {
				flag2 := true
				for k := 0; k < len(slcMap); k++ {
					if !(slcMap[k].Property == resMap[k].Property && slcMap[k].Condition == resMap[k].Condition) {
						flag2 = false
						break
					}
				}
				if flag2 == true {
					flag = false // 存在重复元素，标识为false
					break
				}
			}

		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}
