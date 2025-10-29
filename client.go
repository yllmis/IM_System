package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}

	client.conn = conn
	return client
}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("===功能菜单===")
	fmt.Println("1. 公聊模式")
	fmt.Println("2. 私聊模式")
	fmt.Println("3. 修改用户名")
	fmt.Println("0. 退出登录")
	fmt.Print("请选择(0-3): ")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>请选择合法范围内的数字<<<")
		return false
	}
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}
		switch client.flag {
		case 1:
			fmt.Println("公聊模式选择...")
		case 2:
			fmt.Println("私聊模式选择...")
		case 3:
			fmt.Println("修改用户名选择...")
		}
	}
}

var serverIp string
var serverPort int

// ./server -ip 127.0.0.1 -port 5080
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器ip地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 5080, "设置服务器端口(默认是5080)")
}

func main() {
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>Failed to create client")
		return
	}
	fmt.Println(">>>>>Client connected to server:", client.ServerIp, ":", client.ServerPort)
	client.Run()
}
