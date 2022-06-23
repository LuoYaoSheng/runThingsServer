package aepapi

import (
	"net/http"
	aepsdkcore "run-things-server/extend/aepApis/core"
)

func QueryServiceState(appKey string, appSecret string) (*http.Response, error) {
	path := "/aep_mq_sub/mqStat"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201218144210"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func OpenMqService(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_mq_sub/mqStat"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201217094438"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func CreateTopic(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_mq_sub/topic"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201218142343"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数topicId: 类型long, 参数不可以为空
//  描述:
func QueryTopicInfo(appKey string, appSecret string, topicId string) (*http.Response, error) {
	path := "/aep_mq_sub/topic"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["topicId"] = topicId

	version := "20201218153403"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数topicId: 类型long, 参数不可以为空
//  描述:
func QueryTopicCacheInfo(appKey string, appSecret string, topicId string) (*http.Response, error) {
	path := "/aep_mq_sub/topic/cache"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["topicId"] = topicId

	version := "20201218150354"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

func QueryTopics(appKey string, appSecret string) (*http.Response, error) {
	path := "/aep_mq_sub/topics"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201218153456"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "GET")
}

//参数topicId: 类型long, 参数不可以为空
//  描述:
//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func ChangeTopicInfo(appKey string, appSecret string, topicId string, body string) (*http.Response, error) {
	path := "/aep_mq_sub/topic"

	var headers map[string]string = nil
	var param map[string]string = make(map[string]string)
	param["topicId"] = topicId

	version := "20201218153044"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func QuerySubRules(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_mq_sub/rule"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201218160237"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "POST")
}

//参数body: 类型json, 参数不可以为空
//  描述:body,具体参考平台api说明
func ChangeSubRules(appKey string, appSecret string, body string) (*http.Response, error) {
	path := "/aep_mq_sub/rule"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201218161013"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, body, version, application, key, "PUT")
}

func ClosePushService(appKey string, appSecret string) (*http.Response, error) {
	path := "/aep_mq_sub/mqStat"

	var headers map[string]string = nil
	var param map[string]string = nil
	version := "20201217141937"

	application := appKey
	key := appSecret

	return aepsdkcore.SendAepHttpRequest(path, headers, param, "", version, application, key, "DELETE")
}
