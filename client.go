package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
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

func (client *Client) DealResponse() {
	//一旦client.conn有数据，就copy到stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, client.conn)
	// for {
	// 	buf := make()
	// 	client.conn.Read(buf)
	// 	fmt.Println(string(buf))
	// }
}

func (client *Client) menu() bool {
	var flag int

	fmt.Println("===功能菜单===")
	fmt.Println("1. 公聊模式")
	fmt.Println("2. 私聊模式")
	fmt.Println("3. 修改用户名")
	fmt.Println("0. 退出登录")
	fmt.Println("请选择(0-3): ")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>请选择合法范围内的数字<<<")
		return false
	}
}

func (client *Client) UpdateName() bool {
	fmt.Print(">>>>>请输入用户名: ")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
}

func (client *Client) PublicChat() {
	var chatMsg string

	fmt.Println(">>>>>请输入聊天内容，exit退出。")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn.Write err:", err)
				break
			}

		}
		fmt.Println(">>>>>请输入聊天内容，exit退出。")
		fmt.Scanln(&chatMsg)
	}
	fmt.Println(">>>>>退出公聊模式...")
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}
		switch client.flag {
		case 1:
			client.PublicChat()
			break
		case 2:
			fmt.Println("私聊模式选择...")
			break
		case 3:
			client.UpdateName()
			break
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

	go client.DealResponse()

	fmt.Println(">>>>>Client connected to server:", client.ServerIp, ":", client.ServerPort)
	client.Run()
}
