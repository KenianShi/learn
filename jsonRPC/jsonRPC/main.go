package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
	"net"
	"log"
)

// 需要传输的对象
type RpcObj struct {
	Id   int `json:"id"` // struct标签， 如果指定，jsonrpc包会在序列化json时，将该聚合字段命名为指定的字符串
	Name string `json:"name"`
}

// 需要传输的对象
type ReplyObj struct {
	Ok  bool `json:"ok"`
	Id  int `json:"id"`
	Msg string `json:"msg"`
}

// server端的rpc处理器
type ServerHandler struct {}

// server端暴露的rpc方法
func (serverHandler ServerHandler) GetName(id int, returnObj *RpcObj) error {
	log.Println("server\t-", "recive GetName call, id:", id)
	returnObj.Id = id
	returnObj.Name = "名称1"
	return nil
}

// server端暴露的rpc方法
func (serverHandler ServerHandler) SaveName(rpcObj *RpcObj, returnObj *ReplyObj) error {
	log.Println("server\t-", "recive SaveName call, RpcObj:", rpcObj)
	returnObj.Ok = true
	returnObj.Id = rpcObj.Id
	returnObj.Msg = "存储成功"
	return nil
}

// 开启rpc服务器
func startServer() {
	// 新建Server
	server := rpc.NewServer()

	// 开始监听,使用端口 8888
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("server\t-", "listen error:", err.Error())
	}
	defer listener.Close()

	log.Println("server\t-", "start listion on port 8888")

	// 新建处理器
	serverHandler := ServerHandler{}

	// 注册处理器
	server.Register(serverHandler)

	// 等待并处理链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}

		// 在goroutine中处理请求
		// 绑定rpc的编码器，使用http connection新建一个jsonrpc编码器，并将该编码器绑定给http处理器
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

// 客户端以同步的方式向rpc服务器请求
func callRpcBySynchronous() {
	// 连接至服务器
	client, err := net.DialTimeout("tcp", "localhost:8888", 1000*1000*1000*30) // 30秒超时时间
	if err != nil {
		log.Fatal("client\t-", err.Error())
	}

	defer client.Close()

	// 建立rpc通道
	clientRpc := jsonrpc.NewClient(client)


	// 远程服务器返回的对象
	var rpcObj RpcObj
	log.Println("client\t-", "call GetName method")
	// 请求数据，rpcObj对象会被填充
	clientRpc.Call("ServerHandler.GetName", 1, rpcObj)
	log.Println("client\t-", "recive remote return", rpcObj)


	// 远程返回的对象
	var reply ReplyObj

	// 传给远程服务器的对象参数
	saveObj := RpcObj{2, "对象2"}

	log.Println("client\t-", "call SetName method")
	// 请求数据
	clientRpc.Call("ServerHandler.SaveName", saveObj, reply)

	log.Println("client\t-", "recive remote return", reply)
}

// 客户端以异步的方式向rpc服务器请求
func callRpcByAsynchronous() {
	// 打开链接
	client, err := net.DialTimeout("tcp", "localhost:8888", 1000*1000*1000*30) // 30秒超时时间
	if err != nil {
		log.Fatal("client\t-", err.Error())
	}
	defer client.Close()

	// 建立rpc通道
	clientRpc := jsonrpc.NewClient(client)

	// 用于阻塞主goroutine
	endChan := make(chan int, 15)

	// 15次请求
	for i := 1 ; i <= 15; i++ {

		// 传给远程的对象
		saveObj := RpcObj{i, "对象"}

		log.Println("client\t-", "call SetName method")
		// 异步的请求数据
		divCall := clientRpc.Go("ServerHandler.SaveName", saveObj, ReplyObj{}, nil)

		// 在一个新的goroutine中异步获取远程的返回数据
		go func(num int) {
			reply := <-divCall.Done
			log.Println("client\t-", "recive remote return by Asynchronous", reply.Reply)
			endChan <- num
		}(i)
	}

	// 15个请求全部返回时此方法可以退出了
	for i := 1 ; i <= 15; i++ {
		_ = <-endChan
	}

}

func main() {
	go startServer()
	callRpcBySynchronous()
	callRpcByAsynchronous()
	for{
		;
	}
}