package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const rpcURL = "https://eth-sepolia.g.alchemy.com/v2/qQOTSM74KL3PgdA207aoX"
const privateKeyStr = "dc846d75725e1ca9c22b87a48762e7aecc1598a6397e890b637e3eb95b951bfe"

func main() {
	// 部署合约
	contractAddress := deployContract()
	// 打印合约地址
	fmt.Println(contractAddress)
	// 执行合约
	executeContract(contractAddress)
}

// 部署合约
func deployContract() string {
	fmt.Println("部署合约<-------------------------------->开始")
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

	// 获取gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 获取网络ID
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// 创建交易
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	// 设置非确认交易数量
	auth.Nonce = big.NewInt(int64(nonce))
	// 设置转账金额
	auth.Value = big.NewInt(0) // in wei
	// 设置gas限制
	auth.GasLimit = uint64(300000) // in units
	// 设置gas价格
	auth.GasPrice = gasPrice
	// 设置合约输入
	input := "1.0"
	// 部署合约
	address, tx, instance, err := DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := address.Hex()
	// 打印合约地址
	fmt.Println(contractAddress)
	// 打印交易哈希
	fmt.Println(tx.Hash().Hex())
	// 打印合约实例
	_ = instance

	// 等待部署交易确认
	fmt.Println("等待部署交易确认...")
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status == 0 {
		log.Fatal("部署交易失败")
	}
	fmt.Println("部署交易已确认，区块号:", receipt.BlockNumber)

	fmt.Println("部署合约<-------------------------------->结束")
	return contractAddress
}

// 执行合约
func executeContract(contractAddress string) {
	fmt.Println("执行合约<-------------------------------->开始")
	// 连接到以太坊网络
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal(err)
	}
	// 创建合约实例
	storeContract, err := NewStore(common.HexToAddress(contractAddress), client)
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

	// 获取gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 获取网络ID
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 设置key和value
	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("demo_save_key"))
	copy(value[:], []byte("demo_save_value11111"))

	// 创建交易
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	// 设置非确认交易数量
	opt.Nonce = big.NewInt(int64(nonce))
	// 设置转账金额
	opt.Value = big.NewInt(0) // in wei
	// 设置gas限制
	opt.GasLimit = uint64(300000) // in units
	// 设置gas价格
	opt.GasPrice = gasPrice

	// 设置item
	tx, err := storeContract.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	// 等待交易确认
	fmt.Println("等待交易确认...")
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status == 0 {
		log.Fatal("交易失败")
	}
	fmt.Println("交易已确认，区块号:", receipt.BlockNumber)

	// 创建调用选项
	callOpt := &bind.CallOpts{Context: context.Background()}
	// 获取合约中的value
	valueInContract, err := storeContract.Items(callOpt, key)
	if err != nil {
		log.Fatal(err)
	}
	// 打印合约中的value是否等于原始value
	fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)

	// 打印执行合约结束
	fmt.Println("执行合约<-------------------------------->结束")
}
