package main

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Create an account
	key, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println(err)
		return
	}

	//geth account new
	//geth account list

	//我们知道加密钥匙对由公钥和私钥组成，理论上来讲可能存在两个相同的钥匙对，但这个可能性非常非常非常小。
	//生成一个以太坊公钥有三个步骤：
	//1.随机的生成256bit私钥。随机选取的私钥保证了安全性，只要有足够的随机性，其他人就不可能产生跟你相同的私钥。
	//2. 使用椭圆曲线签名算法elliptic curve cryptography将私钥映射(sha-3)生成公钥。
	////openssl ecparam -name secp256k1 -genkey -noout |openssl ec -text -noout
	//3. 用公钥低位的160bit通过SHA-3加密hash算法计算得到公共地址。
	//eth public addr = sha-3(low 160bit(256bit的私钥))
	//取低160bit意味着一个ethereum账户可以对应于不止一个EC私钥

	// Get the address
	// 生成Ethereum地址 256bit的私钥映射到160bit的公共地址
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	fmt.Printf("ethereum address [%d]:%v\n", len(address), address)
	// address [42]:0x91c0c3e281344d696387Ee2Cd56a9ef198873532

	// Get the private key 256bit的私钥映射到160bit的公共地址
	privateKey := hex.EncodeToString(key.D.Bytes())
	fmt.Printf("EC private key: [%d]%v\n", len(privateKey), privateKey)
	//private key: [64]3bbfab4857df5a142f20f5fd6294e80819e281e0e6bcef0d44904c93f697978a

}
