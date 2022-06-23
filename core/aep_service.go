package service

import (
	"encoding/json"
	"fmt"
	"github.com/LuoYaoSheng/runThingsConfig/config"
	aepapi "github.com/LuoYaoSheng/runThingsServer/extend/aepApis/Aep_device_command"
	aepapi2 "github.com/LuoYaoSheng/runThingsServer/extend/aepApis/Aep_device_management"
	"io/ioutil"

	"strconv"
)

type AepCmdBody struct {
	Content       AepCmdBodyContent `json:"content"`
	DeviceId      string            `json:"deviceId"`
	Operator      string            `json:"operator"`
	ProductId     int64             `json:"productId"`
	Ttl           int               `json:"ttl"`
	DeviceGroupId int               `json:"deviceGroupId"`
	Level         int               `json:"level"`
}

type AepCmdBodyContent struct {
	ServiceIdentifier string `json:"serviceIdentifier"`
	Params            string `json:"params"` // json 字符串
}

type AepCmdRespone struct {
	Code   int                 `json:"code"`
	Msg    string              `json:"msg"`
	Result AepCmdResponeResult `json:"result"`
}

type AepCmdResponeResult struct {
	CommandId     string `json:"commandId"`
	Command       string `json:"command"`
	CommandStatus string `json:"commandStatus"`
	ProductId     int64  `json:"productId"`
	DeviceId      string `json:"deviceId"`
	Imei          string `json:"imei"`
	CreateBy      string `json:"createBy"`
	CreateTime    int64  `json:"createTime"`
	Ttl           int    `json:"ttl"`
}

type AepCreateOther struct {
	AutoObserver int64  `json:"autoObserver"`
	Imsi         string `json:"imsi,omitempty"`
	PskValue     string `json:"pskValue,omitempty"`
}

type AepCreateBody struct {
	DeviceName string         `json:"deviceName"`
	DeviceSn   string         `json:"deviceSn,omitempty"` // sn 和 imei 至少填一个
	Imei       string         `json:"imei,omitempty"`     // sn 和 imei 至少填一个
	Operator   string         `json:"operator"`
	Other      AepCreateOther `json:"other"`
	ProductId  int64          `json:"productId"`
}

type AepListByDeviceIdsBody struct {
	ProductId    int64    `json:"productId"`
	DeviceIdList []string `json:"deviceIdList"`
}

type AepRespone struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
	//Result string `json:"result"`
}

type AepCreateResponeResult struct {
	DeviceSn   string `json:"deviceSn"`
	DeviceId   string `json:"deviceId"`
	DeviceName string `json:"deviceName"`
	TenantId   string `json:"tenantId"`
	ProductId  int64  `json:"productId"`
	Imei       string `json:"imei"`
	Token      string `json:"token"`
}

type AepListByDeviceIdsItem struct {
	ActiveTime      int64  `json:"activeTime"`      //,omitempty"`
	CreateBy        string `json:"createBy"`        //,omitempty"`
	CreateTime      int64  `json:"createTime"`      //,omitempty"`
	DeviceId        string `json:"deviceId"`        //,omitempty"`
	DeviceName      string `json:"deviceName"`      //,omitempty"`
	DeviceSn        string `json:"deviceSn"`        //,omitempty"`
	DeviceStatus    int64  `json:"deviceStatus"`    //,omitempty"`
	FirmwareVersion string `json:"firmwareVersion"` //,omitempty"`
	Imei            string `json:"imei"`            //,omitempty"`
	Imsi            string `json:"imsi"`            //,omitempty"`
	LogoutTime      string `json:"logoutTime"`      //,omitempty"`
	NetStatus       int64  `json:"netStatus"`       //,omitempty"`
	OfflineAt       int64  `json:"offlineAt"`       //,omitempty"`
	OnlineAt        int64  `json:"onlineAt"`        //,omitempty"`
	ProductId       int64  `json:"productId"`       //,omitempty"`
	TenantId        string `json:"tenantId"`        //,omitempty"`
	UpdateBy        string `json:"updateBy"`        //,omitempty"`
	UpdateTime      int64  `json:"updateTime"`      //,omitempty"`
}

type AepListByDeviceIdsRespone struct {
	Code   int                      `json:"code"`
	Msg    string                   `json:"msg"`
	Result []AepListByDeviceIdsItem `json:"result"`
}

// AepCmd 指令下发
func AepCmd(cfg *config.AepConf, sn string, parameter string, serviceIdentifier string, operator string) error {

	bodyContent := AepCmdBodyContent{
		Params:            parameter,
		ServiceIdentifier: serviceIdentifier,
	}
	body := AepCmdBody{
		Content:       bodyContent,
		DeviceId:      sn,
		Operator:      operator,
		ProductId:     cfg.ProductId,
		Ttl:           7200,
		DeviceGroupId: 0,
		Level:         1,
	}
	bodyBytes, err := json.Marshal(body)
	fmt.Println("~~~~~bodyBytes:", string(bodyBytes))

	if err != nil {
		return err
	}

	resp, err1 := aepapi.CreateCommand(cfg.AppKey, cfg.AppSecret, cfg.MasterKey, string(bodyBytes))

	bodyStr, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("--------------bodyStr: ", string(bodyStr))

	return err1
}

// AepCreate 增加设备：没有批量方法
func AepCreate(cfg *config.AepConf, sn, imei, deviceName, operator string) (*AepRespone, error) {

	createOther := AepCreateOther{
		AutoObserver: 1,
	}

	body := AepCreateBody{
		DeviceName: deviceName,
		DeviceSn:   sn,
		Imei:       imei,
		Operator:   operator,
		ProductId:  cfg.ProductId,
		Other:      createOther,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err1 := aepapi2.CreateDevice(cfg.AppKey, cfg.AppSecret, cfg.MasterKey, string(bodyBytes))

	if err1 != nil {
		return nil, err1
	}

	bodyStr, _ := ioutil.ReadAll(resp.Body)

	rsp := AepRespone{}
	err2 := json.Unmarshal(bodyStr, &rsp)
	if err2 != nil {
		return nil, err2
	}

	fmt.Println("--------------bodyStr: ", string(bodyStr))

	return &rsp, nil
}

// AepDeviceList 批量获取设备信息
func AepDeviceList(cfg *config.AepConf, pageNow, pageSize int64, searchValue string) (*AepRespone, error) {

	resp, err := aepapi2.QueryDeviceList(cfg.AppKey, cfg.AppSecret, cfg.MasterKey, strconv.FormatInt(cfg.ProductId, 10), searchValue, strconv.FormatInt(pageNow, 10), strconv.FormatInt(pageSize, 10))
	if err != nil {
		return nil, err
	}

	bodyStr, _ := ioutil.ReadAll(resp.Body)

	rsp := AepRespone{}
	err2 := json.Unmarshal(bodyStr, &rsp)
	if err2 != nil {
		return nil, err2
	}

	fmt.Println("--------------bodyStr: ", string(bodyStr))

	return &rsp, nil
}

func AepListByDeviceIds(cfg *config.AepConf, deviceIdList []string) (*AepListByDeviceIdsRespone, error) {

	body := AepListByDeviceIdsBody{
		ProductId:    cfg.ProductId,
		DeviceIdList: deviceIdList,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := aepapi2.ListDeviceInfo(cfg.AppKey, cfg.AppSecret, cfg.MasterKey, string(bodyBytes))
	if err != nil {
		return nil, err
	}

	bodyStr, _ := ioutil.ReadAll(resp.Body)

	rsp := AepListByDeviceIdsRespone{}
	err2 := json.Unmarshal(bodyStr, &rsp)
	if err2 != nil {
		return nil, err2
	}

	return &rsp, nil
}
