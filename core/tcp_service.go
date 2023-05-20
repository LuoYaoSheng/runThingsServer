package service

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/LuoYaoSheng/runThingsConfig/config"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/LuoYaoSheng/runThingsConfig/model"
	"github.com/LuoYaoSheng/runThingsServer/utils"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/v2/net/gtcp"
)

// TcpReplyFunc tcp回复函数
type TcpReplyFunc func(conn *gtcp.Conn, clientPort string, data []byte)
type TcpResolvingFunc func(ip, clientPort string, data []byte)
type TcpOfflineFunc func(port string) (sq *model.Eq2MqLog)

// TcpServer 因转发需要用到RabbitMQ，所以需先初始化RabbitMQ
func TcpServer(tcpPort, debug int, replyFunc TcpReplyFunc, resolvingFunc TcpResolvingFunc, offlineFunc TcpOfflineFunc) {
	go gtcp.NewServer(":"+strconv.Itoa(tcpPort), func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				if debug > 0 {
					fmt.Println("接收到的数据: ", len(data), hex.EncodeToString(data), conn.RemoteAddr())
				}

				go func() {
					ip, _ := utils.GetLocalIp()
					port := conn.RemoteAddr().String()

					if replyFunc != nil {
						replyFunc(conn, port, data)
					}
					if resolvingFunc != nil {
						resolvingFunc(ip, port, data)
					}
				}()
			} else {
				if debug > 0 {
					fmt.Println("释放端口：", conn.RemoteAddr())
				}
				if offlineFunc != nil {
					offlineFunc(conn.RemoteAddr().String())
				}
			}
			if err != nil {
				break
			}
		}
	}).Run()
}

// TcpOnline_ 上线例子
func TcpOnline_(port string, sn string) (sq *model.Eq2MqLog) {
	return nil
}

// TcpOffline_ 离线例子
func TcpOffline_(port string) (sq *model.Eq2MqLog) {
	return nil
}

// TcpClientSend_ 发送例子
func TcpClientSend_(sn string, snConnMap map[string]*gtcp.Conn, data []byte) {
	conn := snConnMap[sn]
	err := conn.Send(data)
	if err != nil {
		fmt.Println("conn.Send err:", err)
		return
	}
}
