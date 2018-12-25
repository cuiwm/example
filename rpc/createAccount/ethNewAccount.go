package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

// Result is sample

type Result struct {
	result bool
}

func main() {
	//rpc.HandleHTTP()
	//client, er := rpc.Dial("http://192.168.32.146:8545")
	client, er := rpc.Dial("http://127.0.0.1:8545")
	if er != nil {
		fmt.Print("err:", er)
		return
	}

	//	var resp bool
	//	if err := client.Call(&resp, "net_listening"); err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println("resp:", resp)
	//
	//请求：
	//-->{"id":1, "jsonrpc":"2.0","method":"eth_sendConsensusScore",
	//"params":[{"0x26f3ec8ac34476cea0a04175d0a891fbdb6d7ab0",
	//"password"，
	//"0x1",
	//"timestamp",
	//"pieceid",
	//"sign"}]}
	//param1: 上传者帐号地址
	//param2: 密码
	//param3: 流量积分数量,即一个内容碎片的积分数
	//param4: Unix时间戳
	//param5: 流量碎片id
	//param6: 流量signHash
	//sign生成：
	//md5(param4 + param3 + （param5）的内容) = signHash
	//	//curl -H "Content-Type:application/json"   -X POST --data
	//'{"jsonrpc":"2.0","method":"eth_sendConsensusScore",
	//	"params":[{"0x3915aa98fe6ca5d5968c5ba01e24d1d8ba341d92", "123456","0x100","12343214","998","900150983cd24fb0d6963f7d28e17f72"}],"id":67}' 192.168.96.248:8545

	type NewAccountArgs struct {
		Password string `json:"passwd"`
	}
	var req string = "12343252"
	var resp string
	if err := client.Call(&resp, "personal_newAccount", req); err != nil {
		//if err := client.Call(&resp, "personal_newAccount", &NewAccountArgs{"134234213"}); err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Printf("resp:%s\n", resp)

}
