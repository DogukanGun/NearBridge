package contracts

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"strings"
)

func CreateFunctionRequirementsForBridge(clientUrl string, contractAddress string, privateKey string) (error, *Bridge, common.Address, *bind.TransactOpts, *ethclient.Client) {
	err, client, _publicAddress, res := CreateFunctionRequirementsForControllers(
		clientUrl,
		"Bridge.abi",
		contractAddress,
		privateKey)

	address := common.HexToAddress(contractAddress)
	contractInstance, err := NewBridge(address, client)
	return err, contractInstance, _publicAddress, res, client
}

func CreateFunctionRequirementsForControllers(clientUrl string, walletAbiName string, contractAddress string, privateKey string) (error, *ethclient.Client, common.Address, *bind.TransactOpts) {
	client, err := ethclient.Dial(clientUrl)
	if err != nil {
		// Handle error
	}

	address := common.HexToAddress(contractAddress)
	abiFile, err := ioutil.ReadFile(walletAbiName)
	contractAbi, err := abi.JSON(strings.NewReader(string(abiFile)))
	if err != nil {
		// Handle error
		fmt.Println(err)
	}

	fmt.Println(address)
	fmt.Println(client)
	fmt.Println(contractAbi)

	if err != nil {
		// Handle error
	}

	_privateKey, _, _publicAddress, _ := GenerateKeypairFromPrivateKeyHex(privateKey)
	res, _ := BuildTransactionOptions(client, _privateKey, 300000)
	fmt.Println(res)
	return err, client, _publicAddress, res
}

func BuildTransactionOptions(client *ethclient.Client, prvKey *ecdsa.PrivateKey, gasLimit uint64) (*bind.TransactOpts, error) {

	// Retrieve the chainID
	chainID, cIDErr := client.ChainID(context.Background())

	if cIDErr != nil {
		return nil, cIDErr
	}

	// Retrieve suggested gas price
	gasPrice, gErr := client.SuggestGasPrice(context.Background())

	if gErr != nil {
		return nil, gErr
	}

	// Create empty options object
	txOptions, txOptErr := bind.NewKeyedTransactorWithChainID(prvKey, chainID)

	if txOptErr != nil {
		return nil, txOptErr
	}

	txOptions.Value = big.NewInt(0) // in wei
	txOptions.GasLimit = gasLimit   // in units
	txOptions.GasPrice = gasPrice

	return txOptions, nil
}

func GenerateKeypairFromPrivateKeyHex(privateKeyHex string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, common.Address, error) {

	log.Println("Generating the keypair...")

	// If hex string has "0x" at the start discard it
	if privateKeyHex[:2] == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}

	// Convert hex key to a private key object
	privateKey, privateKeyErr := crypto.HexToECDSA(privateKeyHex)

	if privateKeyErr != nil {
		return nil, nil, common.Address{}, privateKeyErr
	}

	// Generate the public key using the private key
	publicKey := privateKey.Public()

	// Cast crypto.Publickey object to the ecdsa.PublicKey object
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		return nil, nil, common.Address{}, errors.New("error casting public key to ECDSA")
	}

	// Convert publickey to a address
	pubkeyAsAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, publicKeyECDSA, pubkeyAsAddress, nil
}
