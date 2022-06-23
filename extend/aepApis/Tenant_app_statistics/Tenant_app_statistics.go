package aepapi

import (
	aepsdkcore "github.com/LuoYaoSheng/runThingsServer/extend/aepApis/core"
	"net/http"
)

func QueryTenantApiMonthlyCount(appKey string, appSecret string) (*http.Response, error) {
	path := "/tenant_app_statistics/queryTenantApiMonthlyCount"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201225122609"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

func QueryTenantAppCount(appKey string, appSecret string) (*http.Response, error) {
	path := "/tenant_app_statistics/queryTenantAppCount"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201225122611"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数dateType: 类型String, 参数不可以为空
//  描述:日期格式 m：月  d：日
//参数dataType: 类型String, 参数不可以为空
//  描述:数据格式 1：api调用量分析
func QueryTenantApiTrend(appKey string, appSecret string, dateType string, dataType string) (*http.Response, error) {
	path := "/tenant_app_statistics/queryTenantApiTrend"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["dateType"] = dateType
	param["dataType"] = dataType

	version := "20201225122606"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}
