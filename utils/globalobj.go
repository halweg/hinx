package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"zinx/izface"
)

/**
	在这里存储一些全局配置
 */

type  GlobalObj struct {

	//server
	TcpServer izface.IServer
	Host string
	TcpPort string
	Name string

	//zinx
	Version string
	MaxConn int
	MaxPackageSize uint32

}

func (g *GlobalObj) Reload () {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		fmt.Println("配置文件加载失败！")
		panic(err)
	}
	fmt.Println(string(data))
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

var GlobalObject *GlobalObj

func init() {
	//如果配置没有加载，这里设置默认的值
	GlobalObject = &GlobalObj{
		Name: "Zinx server app",
		Version: "v0.4",
		Host: "0.0.0.0",
		TcpPort: "8999",
		MaxConn: 1000,
		MaxPackageSize: 4096,
	}

	//GlobalObject.Reload()

}