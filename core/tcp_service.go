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

var snConnMap map[string]*gtcp.Conn

// TcpReplyFunc tcp回复函数
type TcpReplyFunc func(conn *gtcp.Conn, data []byte)
type TcpResolvingFunc func(ip, port string, data []byte)

// TcpServer 因转发需要用到RabbitMQ，所以需先初始化RabbitMQ
func TcpServer(tcpPort int, replyFunc TcpReplyFunc, resolvingFunc TcpResolvingFunc) {

	if LogsMQ == nil {
		panic("请先初始化RabbitMQ")
	}

	snConnMap = make(map[string]*gtcp.Conn)

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

func TcpReply(conn *gtcp.Conn, data []byte) {

	//sn := hex.EncodeToString(data[12:20])
	//snConnMap[sn] = conn
	// 得找到sn

	if len(data) == 12 {
		r, _ := hex.DecodeString("404000000000000000002323")
		if err := conn.Send(r); err != nil {
			fmt.Println(err)
		}
	} else {
		if data[4] == 0x02 && data[5] == 0x03 {
			//data2 := [30]byte{0x00}
			var data2 []byte
			data2 = make([]byte, 30)

			for i := 0; i < 12; i++ {
				data2[i] = data[i]
			}
			for i := 12; i < 18; i++ {
				data2[i] = data[i+6]
			}
			for i := 18; i < 24; i++ {
				data2[i] = data[i-6]
			}
			data2[24] = 0x00
			data2[25] = 0x00
			data2[26] = 0x03
			data2[27] = 0x00
			data2[28] = 0x23
			data2[29] = 0x23

			// 修复数据
			//data2[2] = 0x00
			data2[27] = byte(utils.CheckNum(data2[2:27]))
			//fmt.Println("测验数据: ", hex.EncodeToString(data2))
			//fmt.Printf("------ 数据：%x \n", data2[27])
			fmt.Println("-------- 应答数据: ", hex.EncodeToString(data2))
			if err := conn.Send(data2); err != nil {
				fmt.Println("conn.Send err:", err)
			}
		}
	}
}

func TcpClientSend(sn string, data []byte) {
	conn := snConnMap[sn]
	err := conn.Send(data)
	if err != nil {
		fmt.Println("conn.Send err:", err)
		return
	}
}
