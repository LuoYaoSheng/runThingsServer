package aepapi

import (
	aepsdkcore "github.com/LuoYaoSheng/runThingsServer/extend/aepApis/core"
	"net/http"
)

func QueryTenantDeviceCount(appKey string, appSecret string) (*http.Response, error) {
	path := "/tenant_device_statistics/queryTenantDeviceCount"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201225122555"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数dateType: 类型String, 参数不可以为空
//  描述:时间类型：d:天；m:月
//参数type1: 类型String, 参数不可以为空
//  描述:数据类型：1：设备总数量，激活数，活跃数；3：设备活跃数，活跃率
func QueryTenantDeviceTrend(appKey string, appSecret string, dateType string, type1 string) (*http.Response, error) {
	path := "/tenant_device_statistics/queryTenantDeviceTrend"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["dateType"] = dateType
	param["type1"] = type1

	version := "20201225122550"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

func QueryTenantDeviceActiveCount(appKey string, appSecret string) (*http.Response, error) {
	path := "/tenant_device_statistics/queryTenantDeviceActiveCount"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201225122558"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}
