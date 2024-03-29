package service

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

var influxClient client.Client

func GetClient(addr, username, password, database, precision string) client.Client {
	if influxClient == nil {
		var err error
		influxClient, err = client.NewHTTPClient(client.HTTPConfig{
			Addr:     addr,
			Username: username,
			Password: password,
		})
		if err != nil {
			log.Fatalln(err)
			return nil
		}

		// 检测是否有库，没有则创建
		createDbSQL := client.NewQuery(fmt.Sprintf("CREATE DATABASE %s", database), "", "")
		if _, err1 := influxClient.Query(createDbSQL); err1 != nil {
			influxClient = nil
			log.Println(err1)
			return nil
		}
		// 过期策略
		createRPSQL := client.NewQuery(fmt.Sprintf("CREATE RETENTION POLICY default ON %s DURATION 360d REPLICATION 1 DEFAULT", database), database, precision)
		if _, err2 := influxClient.Query(createRPSQL); err2 != nil {
			influxClient = nil
			log.Println(err2)
			return nil
		}
	}
	return influxClient
}

// query
func queryDB(cli client.Client, database, cmd string) (res []client.Result, err1 error) {
	q := client.Query{
		Command:  cmd,
		Database: database,
	}
	if response, err := cli.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

// insert
func writesPoints(cli client.Client, points []*client.Point, database, precision string) error {

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database,
		Precision: precision, //精度，默认ns
	})
	if err != nil {
		log.Println(err)
		return err
	}
	bp.AddPoints(points)
	err = cli.Write(bp)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func WirteInflux(sn string, productKey string, status int64, fields map[string]interface{}, database, prefix, precision string) (string, error) {

	if influxClient == nil {
		return "初始化失败", errors.New("influxdb未初始化")
	}
	// 生成 tags
	tagStatus := `0`
	if status > 0 {
		tagStatus = strconv.FormatInt(status, 10)
	}
	tags := map[string]string{
		"sn":         sn,
		"status":     tagStatus,
		"productKey": productKey,
	}

	// 生成节点
	pt, err := client.NewPoint(prefix+tags["sn"], tags, fields, time.Now())
	if err != nil {
		log.Println(err)
		return "", err
	}
	points := []*client.Point{pt} // 后期会做批量处理，当前先保留数组模式
	err = writesPoints(influxClient, points, database, precision)
	return "", err
}

// SelectInflux sql是赛选条件，例：`LIMIT 10`
func SelectInflux(sn, sql string, database, prefix string) (res []client.Result, err1 error) {
	//if influxClient == nil {
	//	return nil, errors.New("influxdb未初始化")
	//}
	//measure := prefix + sn
	//qs := fmt.Sprintf("SELECT * FROM %s %s", measure, sql)
	//return queryDB(influxClient, database, qs)
	return SelectFieldsInflux(`*`, sn, sql, database, prefix)
}

// SelectFieldsInflux fields,例：`a,b` sql是赛选条件，例：`LIMIT 10`
func SelectFieldsInflux(fields, sn, sql string, database, prefix string) (res []client.Result, err1 error) {
	if influxClient == nil {
		return nil, errors.New("influxdb未初始化")
	}
	measure := prefix + sn
	qs := fmt.Sprintf("SELECT %s FROM %s %s", fields, measure, sql)
	return queryDB(influxClient, database, qs)
}
