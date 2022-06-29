package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//创建远程调用函数，函数一般是放在结构体里面
type Goods struct {
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

func (this Goods) AddGoods(req AddGoodsReq, res *AddGoodsRes) error {
	//1、执行增加		模拟
	fmt.Println(req)
	//2、返回增加结果
	*res = AddGoodsRes{
		Success: true,
		Massage: "增加数据成功",
	}
	return nil
}
func main() {
	//1、注册RPC服务
	err1 := rpc.RegisterName("goods", new(Goods))
	if err1 != nil {
		fmt.Println(err1)
	}
	//2、监听端口
	listener, err2 := net.Listen("tcp", "127.0.0.1:8020")
	if err2 != nil {
		fmt.Println(err2)
	}
	//3、关闭监听
	defer listener.Close()
	for {
		//4、监听客户端连接
		fmt.Println("开始监听")
		conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println(err3)
		}
		//5、绑定服务
		//rpc.ServeConn(conn)
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

/*
jsonrpc和默认rpc的区别：
以前rpc.ServeConn(conn)绑定服务
jsonrpc中通过rpc.ServeCodec(jsonrpc.NewServerCodec(conn))绑定服务
*/
