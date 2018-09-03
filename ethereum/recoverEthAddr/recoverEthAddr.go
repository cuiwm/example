package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

//In cryptography, the Elliptic Curve Digital Signature Algorithm (ECDSA) offers a variant of the Digital Signature Algorithm (DSA) which uses elliptic curve cryptography
func main() {
	privKey := "3bbfab4857df5a142f20f5fd6294e80819e281e0e6bcef0d44904c93f697978a"
	expect := "0x91c0c3e281344d696387Ee2Cd56a9ef198873532"
	key, _ := crypto.HexToECDSA(privKey)
	fmt.Println("ecdsa key:", key)
	msg := crypto.Keccak256([]byte("foo"))
	fmt.Println("len:", len(msg), " msg: ", msg)

	sig, _ := crypto.Sign(msg, key)
	//#利用私钥对msg进行签名，同样可以通过签名的msg来恢复公钥，从而恢复地址
	fmt.Println("sig:", sig)
	//1. 利用私钥签名，来推断出私钥对应的公钥
	// crypto.Ecrecover
	recoveredPubBytes, _ := crypto.Ecrecover(msg, sig)
	fmt.Printf("recovered pub bytes:%#x\n", recoveredPubBytes)
	ecdsaPubKey := crypto.ToECDSAPub(recoveredPubBytes)
	fmt.Printf("ecdsa PubKey: %#x\n", ecdsaPubKey)
	recoveredEthAddr := crypto.PubkeyToAddress(*ecdsaPubKey)
	//	// address [42]:0x91c0c3e281344d696387Ee2Cd56a9ef198873532
	fmt.Printf("expect :%s %t\n", recoveredEthAddr.String(), recoveredEthAddr.String() == expect)
	//#common.Address格式的地址

	//2.crypto.SigToPub
	recoveredPub2, _ := crypto.SigToPub(msg, sig)
	recoveredEthAddr2 := crypto.PubkeyToAddress(*recoveredPub2)
	fmt.Printf("expect2:%s %t\n", recoveredEthAddr2.String(), recoveredEthAddr2.String() == expect)

}
