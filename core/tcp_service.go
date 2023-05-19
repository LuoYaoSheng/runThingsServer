package service

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/LuoYaoSheng/runThingsConfig/config"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/LuoYaoSheng/runThingsConfig/model"
	service "github.com/LuoYaoSheng/runThingsServer/core"
	"github.com/LuoYaoSheng/runThingsServer/utils"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/v2/net/gtcp"
)

var SNConnMap map[string]*gtcp.Conn

// TcpReplyFunc tcp回复函数
type TcpReplyFunc func(conn *gtcp.Conn, clientPort string,data []byte)
type TcpResolvingFunc func(ip, clientPort string, data []byte)

// TcpServer 因转发需要用到RabbitMQ，所以需先初始化RabbitMQ
func TcpServer(tcpPort int, replyFunc TcpReplyFunc, resolvingFunc TcpResolvingFunc) {

	if LogsMQ == nil {
		panic("请先初始化RabbitMQ")
	}

	SNConnMap = make(map[string]*gtcp.Conn)

	go gtcp.NewServer(":"+strconv.Itoa(tcpPort), func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				fmt.Println("接收到的数据: ", len(data), hex.EncodeToString(data), conn.RemoteAddr())
				go func() {
					if replyFunc != nil {
						replyFunc(conn, data)
					}
					if resolvingFunc != nil {
						ip, _ := utils.GetLocalIp()
						port := conn.RemoteAddr().String()
						resolvingFunc(ip, port, data)
					}
				}()
			} else {
				fmt.Println("释放端口：", conn.RemoteAddr())
				OfflineTcp(conn.RemoteAddr().String())
			}
			if err != nil {
				break
			}
		}
	}).Run()
}

func OnlineTcp(port string, sn string) {

	portKey := "tcp" + port

	_, err := service.GetRdValue(portKey)
	if err == redis.Nil {
		// 存储设备 & 设备上线
		err4 := service.SetRdValue(portKey, sn)
		if err4 != nil {
			g.Log().Error(bgContext, err4)
			return
		}
		sq := &model.Eq2MqLog{
			Sn:     sn,
			Status: config.EqStatusOnline,
			Title:  "设备上线",
		}
		go Rbmq.LogToMQ(sq)
	} else if err != nil {
		g.Log().Error(bgContext, err)
		return
	}
}

func OfflineTcp(port string) {
	portKey := "tcp" + port

	sn, err := service.GetRdValue(portKey)
	if err != nil {
		g.Log().Error(bgContext, err)
		return
	}
	sq := &model.Eq2MqLog{
		Sn:     sn,
		Status: config.EqStatusOnline,
		Title:  "设备离线",
	}
	go Rbmq.LogToMQ(sq)
}

func TcpClientSend(sn string, data []byte) {
	conn := SNConnMap[sn]
	err := conn.Send(data)
	if err != nil {
		fmt.Println("conn.Send err:", err)
		return
	}
}
