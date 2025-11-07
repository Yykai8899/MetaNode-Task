package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

const privateKeyStr = "dc846d75725e1ca9c22b87a48762e7aecc1598a6397e890b637e3eb95b951bfe"
const rpcURL = "https://eth-sepolia.g.alchemy.com/v2/qQOTSM74KL3PgdA207aoX"

func main() {
	// 获取区块信息
	getBlockInfo()

	// 生成钱包
	toAddress := generateWallet()

	//转账交易
	sendTransaction(toAddress)
}

// 获取区块信息
func getBlockInfo() {
	fmt.Println("获取区块信息<-------------------------------->开始")
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(5671744)

	// 获取区块头信息
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	// 打印区块头信息
	fmt.Println(header.Number.Uint64()) // 5671744
	// 打印区块时间
	fmt.Println(header.Time) // 1712798400
	// 打印区块难度
	fmt.Println(header.Difficulty.Uint64()) // 0
	// 打印区块哈希
	fmt.Println(header.Hash().Hex()) // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	// 获取区块信息
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	// 打印区块信息
	fmt.Println(block.Number().Uint64()) // 5671744
	// 打印区块时间
	fmt.Println(block.Time()) // 1712798400
	// 打印区块难度
	fmt.Println(block.Difficulty().Uint64()) // 0
	// 打印区块哈希
	fmt.Println(block.Hash().Hex()) // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	// 打印区块交易数量
	fmt.Println(len(block.Transactions())) // 70
	// 获取区块交易数量
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	// 打印区块交易数量
	fmt.Println(count) // 70
	fmt.Println("获取区块信息<-------------------------------->结束")
}

// 生成钱包
func generateWallet() string {
	fmt.Println("生成钱包<-------------------------------->开始")
	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// 获取私钥
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// 打印私钥
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 去掉'0x'
	// 获取公钥
	publicKey := privateKey.Public()
	// 断言公钥类型
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 获取公钥字节
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// 打印公钥字节
	fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'
	// 获取地址
	walletAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	// 打印地址
	fmt.Println(walletAddress)
	// 获取哈希
	hash := sha3.NewLegacyKeccak256()
	// 写入公钥字节
	hash.Write(publicKeyBytes[1:])
	// 打印哈希
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	// 钱包地址：原长32位，截去12位，保留后20位
	address := hexutil.Encode(hash.Sum(nil)[12:])
	// 打印钱包地址
	fmt.Println(address) // 原长32位，截去12位，保留后20位
	fmt.Println("生成钱包<-------------------------------->结束")
	return address
}

func sendTransaction(toAddressStr string) {
	fmt.Println("发送交易<-------------------------------->开始")
	// 连接到以太坊网络
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}

	// 转换私钥
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		log.Fatal(err)
	}

	// 获取公钥
	publicKey := privateKey.Public()
	// 断言公钥类型
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 获取地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// 获取非确认交易数量
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 设置转账金额
	value := big.NewInt(1000000000000000000) // in wei (0.001 eth)
	// 设置gas限制
	gasLimit := uint64(21000) // in units
	// 获取gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 转换目标地址
	toAddress := common.HexToAddress(toAddressStr)
	// 设置数据
	var data []byte
	// 创建交易
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// 获取网络ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	// 打印交易哈希
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	fmt.Println("发送交易<-------------------------------->结束")
}
