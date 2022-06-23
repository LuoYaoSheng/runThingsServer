package aepapi

import (
	aepsdkcore "github.com/LuoYaoSheng/runThingsServer/extend/aepApis/core"
	"net/http"
)

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func DeleteEdgeInstanceDevice(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_edge_gateway/instance/devices"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201226000026"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数gatewayDeviceId: 类型String, 参数不可以为空
//  描述:
//参数pageNow: 类型long, 参数可以为空
//  描述:
//参数pageSize: 类型long, 参数可以为空
//  描述:
func QueryEdgeInstanceDevice(appKey string, appSecret string, gatewayDeviceId string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_edge_gateway/instance/devices"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["gatewayDeviceId"] = gatewayDeviceId
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20201226000022"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func CreateEdgeInstance(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_edge_gateway/instance"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201226000017"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func EdgeInstanceDeploy(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_edge_gateway/instance/settings"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201226000010"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数id: 类型long, 参数不可以为空
//  描述:
func DeleteEdgeInstance(appKey string, appSecret string, id string) (*http.Response, error) {
	path := "/aep_edge_gateway/instance"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["id"] = id

	version := "20201225235957"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "DELETE")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func AddEdgeInstanceDevice(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_edge_gateway/instance/device"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201226000004"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func AddEdgeInstanceDrive(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_edge_gateway/instance/drive"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201225235952"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}
