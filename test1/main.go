package main

import (
	"context"
	"flag"
	"github.com/buidl-test/test1/bindings"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

type Event struct {
	Type   string
	TxHash string
	Body   interface{}
}

type TransferEvent struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

func main() {
	addr := flag.String("address", "", "Contract address")
	flag.Parse()

	client, err := connectEthClient()
	if err != nil {
		panic(err)
	}

	contractAddress := *addr //"0xa15D3cBee49516Ef2FD7910AC72C71D74da02710"
	session, err := bindings.NewMyToken(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal("##### failed connect MyToken contract: ", err)
	}

	// トークン情報取得
	tokenInfo(session)

	eventChan := make(chan *Event)
	// イベント監視
	if err := watchEvent(client, contractAddress, eventChan); err != nil {
		log.Println("##### failed watch for event: ", err)
	}

	for {
		event := <-eventChan
		txhash := event.TxHash

		switch event.Type {
		case "Transfer":
			from := event.Body.(TransferEvent).From.Hex()
			to := event.Body.(TransferEvent).To.Hex()
			value := event.Body.(TransferEvent).Value.Uint64()
			log.Println("##### Transfer event info")
			log.Println("     ***** transaction: ", txhash)
			log.Println("     ***** from: ", from)
			log.Println("     ***** to: ", to)
			log.Println("     ***** value: ", value)
		}
	}
}

func connectEthClient() (*ethclient.Client, error) {
	url := "ws://127.0.0.1:8546"
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Println("##### failed connect: ", err)
		return nil, err
	}

	return client, nil
}

func tokenInfo(session *bindings.MyToken) {
	name, _ := session.Name(nil)
	symbol, _ := session.Symbol(nil)
	totalSupply, _ := session.TotalSupply(nil)
	log.Println("##### Token info")
	log.Println("     ***** name: ", name)
	log.Println("     ***** symbol:", symbol)
	log.Println("     ***** totalSupply: ", totalSupply)
}

func watchEvent(client *ethclient.Client, address string, eventChan chan *Event) error {
	topics := map[string]common.Hash{
		"Transfer": crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")),
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(address)},
		Topics: [][]common.Hash{{
			topics["Transfer"],
		}},
	}

	logChan := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logChan)
	if err != nil {
		log.Fatal(err)
	}

	myTokenAbi, err := abi.JSON(strings.NewReader(string(bindings.MyTokenABI)))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Println(err)
				eventChan <- &Event{
					Type:   "Disconnect",
					TxHash: "",
					Body:   "",
				}
				return
			case vLog := <-logChan:
				switch vLog.Topics[0] {
				case topics["Transfer"]:
					var transferEvent TransferEvent
					if err := myTokenAbi.Unpack(&transferEvent, "Transfer", vLog.Data); err != nil {
						log.Println("##### failed to unpack")
						continue
					}

					transferEvent.From = common.BytesToAddress(vLog.Topics[1].Bytes())
					transferEvent.To = common.BytesToAddress(vLog.Topics[2].Bytes())

					eventChan <- &Event{
						Type:   "Transfer",
						TxHash: vLog.TxHash.Hex(),
						Body:   transferEvent,
					}
					break
				}
			default:
			}
		}
	}()

	return nil
}
