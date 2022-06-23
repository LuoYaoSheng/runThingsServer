package service

import (
	"encoding/json"
	"fmt"
	"github.com/LuoYaoSheng/runThingsServer/extend/oneNet"
)

type OneNetDeviceReq struct {
	Ack      bool   `json:"ack"`
	ActTime  string `json:"act_time"`
	Area     string `json:"area"`
	AuthInfo struct {
		Field1 string `json:"869334059269368"`
	} `json:"auth_info"`
	CreateTime  string `json:"create_time"`
	Datastreams []struct {
		CreateTime string `json:"create_time"`
		Id         string `json:"id"`
		Uuid       string `json:"uuid"`
	} `json:"datastreams"`
	Desc     string   `json:"desc"`
	Fversion string   `json:"fversion"`
	Id       string   `json:"id"`
	Imsi     string   `json:"imsi"`
	ImsiMt   string   `json:"imsi_mt"`
	ImsiOld  []string `json:"imsi_old"`
	LastCt   string   `json:"last_ct"`
	Location struct {
		Lat int `json:"lat"`
		Lon int `json:"lon"`
	} `json:"location"`
	Obsv            bool          `json:"obsv"`
	ObsvSt          bool          `json:"obsv_st"`
	Online          bool          `json:"online"`
	Private         bool          `json:"private"`
	Protocol        string        `json:"protocol"`
	RgId            string        `json:"rg_id"`
	SoftwareVersion string        `json:"software_version"`
	Tags            []interface{} `json:"tags"`
	Title           string        `json:"title"`
}

type OneNetDeviceStatusRsp struct {
	Id     string `json:"id"`
	Online bool   `json:"online"`
	Title  string `json:"title"`
}

type OneNetDeviceStatusRspList struct {
	Devices    []*OneNetDeviceStatusRsp `json:"devices"`
	TotalCount int64                    `json:"total_Count"`
}

func OneNetDevice(deviceId int, apiKey string) bool {
	//apiKey := "kE6HpGUkj43NZ6fD=kUTzM=3INE="
	on := oneNet.NewOneNet(apiKey)
	b, str := on.Device(deviceId)
	if b {
		m := OneNetDeviceReq{}
		err := json.Unmarshal([]byte(*str), &m)
		if err != nil {
			fmt.Println("Umarshal failed:", err)
		} else {
			return m.Online
		}
	}
	return false
}
