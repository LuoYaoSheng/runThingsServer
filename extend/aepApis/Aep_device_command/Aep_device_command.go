package aepapi

import (
	aepsdkcore "iot/common/aepApis/core"
	"net/http"
)

// CreateCommand 下发指令
func CreateCommand(appKey string, appSecret string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_device_command/command"

	var headers = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20190712225145"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

// QueryCommandList 查询记录
func QueryCommandList(appKey string, appSecret string, MasterKey string, productId string, deviceId string, startTime string, endTime string, pageNow string, pageSize string) (*http.Response, error) {
	path := "/aep_device_command/commands"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["productId"] = productId
	param["deviceId"] = deviceId
	param["startTime"] = startTime
	param["endTime"] = endTime
	param["pageNow"] = pageNow
	param["pageSize"] = pageSize

	version := "20200814163736"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

// QueryCommand 下发指令
func QueryCommand(appKey string, appSecret string, MasterKey string, commandId string, productId string, deviceId string) (*http.Response, error) {
	path := "/aep_device_command/command"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = make(map[string]string)
	param["commandId"] = commandId
	param["productId"] = productId
	param["deviceId"] = deviceId

	version := "20190712225241"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

// CancelCommand 取消下发指令
func CancelCommand(appKey string, appSecret string, MasterKey string, body string) (*http.Response, error) {
	path := "/aep_device_command/cancelCommand"

	var headers map[string]string = make(map[string]string)
	headers["MasterKey"] = MasterKey

	var param map[string]string = nil
	version := "20190615023142"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}
