package main

import (
	"fmt"
	"net"
	"net/rpc"
)

//定义一个远程调用的方法
type Hello struct {
}

/*
1、方法只能有两个可序列化的参数，其中第二个参数是指针类型
	req 表示获取客户端传过来的数据
	rsq 表示给给客户端返回数据
2、方法要返回一个error类型，同时必须是公开方法
3、req和rsq的类型不能是：channel（通道）、func（函数）均不能进行序列化
*/

func (this Hello) SayHello(req string, res *string) error {
	*res = "你好" + req
	return nil
}
func main() {
	//1、注册RPC服务
	err1 := rpc.RegisterName("hello", new(Hello))
	if err1 != nil {
		fmt.Println(err1)
	}
	//2、监听端口
	listener, err2 := net.Listen("tcp", "127.0.0.1:8080")
	if err2 != nil {
		fmt.Println(err2)
	}
	//应用退出的时候关闭监听端口
	defer listener.Close()
	for {
		fmt.Println("开始建立连接")
		//4、建立连接
		conn, err3 := listener.Accept()
		if err3 != nil {
			fmt.Println(err3)
		}
		//5、绑定服务
		rpc.ServeConn(conn)
	}

}
