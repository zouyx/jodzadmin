package zk

import (
	"github.com/astaxie/beego"
	"github.com/samuel/go-zookeeper/zk"
	"strings"
	"time"
)

var (
	conn *zk.Conn
)

func init() {
	//fist establishes connection
	connect()
}

//do connect with app config
func connect() {
	var err error
	conn, _, err = zk.Connect(getZkIps(), time.Second) //*10)
	if err != nil {
		panic(err)
	}
}

func getZkIps() []string {
	zk := beego.AppConfig.String("zk")

	return cutString(zk)
}

func cutString(str string) []string {
	if str == "" {
		return []string{}
	}

	return strings.Split(str, ",")
}

// get zk connection
func GetConn() *zk.Conn {
	return conn
}
