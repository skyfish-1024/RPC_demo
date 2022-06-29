package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	Hello()
	AddGoods()
}
func Hello() {
	//1、用rpc.Dial()和rpc微服务建立连接
	conn, err1 := rpc.Dial("tcp", "127.0.0.1:8080")
	if err1 != nil {
		fmt.Println(err1)
	}
	//2、当客户端退出关闭连接
	defer conn.Close()
	//3、调用远程函数
	var reply string
	/*
		1、第一个参数	hello.SayHello		hello表示服务名称	SayHello方法名称
		2、第二个参数	给服务端的req传递参数
		3、第三个参数	需要传入地址		获取微服务端返回的数据
	*/
	err2 := conn.Call("hello.SayHello", "我是客户端", &reply)
	if err2 != nil {
		fmt.Println(err2)
	}
	//4、获取微服务返回的数据
	fmt.Println("hello reply:", reply)
}

type AddGoodsReq struct {
	Id      int
	Title   string
	Price   float32
	Content string
}
type AddGoodsRes struct {
	Success bool
	Massage string
}

func AddGoods() {
	//1、用net.Dial()和rpc微服务建立连接
	conn, err1 := net.Dial("tcp", "127.0.0.1:8020")
	if err1 != nil {
		fmt.Println(err1)
	}
	//2、当客户端退出关闭连接
	defer conn.Close()

	//建立基于json编码的rpc服务
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	//3、调用远程函数
	var reply AddGoodsRes
	/*
		1、第一个参数	hello.SayHello		hello表示服务名称	SayHello方法名称
		2、第二个参数	给服务端的req传递参数
		3、第三个参数	需要传入地址		获取微服务端返回的数据
	*/
	err2 := client.Call("goods.AddGoods", AddGoodsReq{
		Id:      10,
		Title:   "我是商品",
		Price:   13,
		Content: "商品详情",
	}, &reply)
	if err2 != nil {
		fmt.Println(err2)
	}
	//4、获取微服务返回的数据
	fmt.Println("Goods reply:", reply)
	fmt.Printf("%#v", reply)

	/*
		把默认的rpc改为jsonrpc
		1、rpc.Dial需要调换成net.Dail
		2、增加建立基于json编码的rpc服务	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
		3、conn.Call需要改为clien.Call
	*/
}
