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
	client, er := rpc.Dial("http://192.168.96.248:8545")
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
	//2.2. 上传流量积分(区块链提供接口服务)
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

	type SendUSArgs struct {
		From       string `json:"from"`
		Passphrase string `json:"passphrase"`
		Score      string `json:"score"`
		Timestamp  string `json:"timestamp"`
		Pieceid    string `json:"pieceid"`
		Sign       string `json:"sign"`
		//Nonce      *hexutil.Uint64 `json:"nonce"`
	}
	var resp string
	if err := client.Call(&resp, "eth_sendConsensusScore", &SendUSArgs{"0x3915aa98fe6ca5d5968c5ba01e24d1d8ba341d92", "123456", "0x100", "123445", "998", "900150983cd24fb0d6963f7d28e17f72"}); err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("ok:", resp)

}
