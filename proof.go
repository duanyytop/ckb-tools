package main

import (
	"context"
	"fmt"

	"github.com/nervosnetwork/ckb-sdk-go/rpc"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

const (
	fromBlock = uint64(5000000)
	toBlock   = uint64(7000000)
)

func main() {
	client, err := rpc.Dial("http://127.0.0.1:8114")
	if err != nil {
		return
	}
	txHashes := []string{"0xb4b691a3ee53b51acd4e8b40ea8bc2fecd3943fa069f2fe7acfd0a0d83c24fe8"}
	blockHash := types.HexToHash("0xef7e8e2159859ec42153e2468ddbfd3890115237366174cc9e7e39087a953f2c")
	txProof, err := client.GetTransactionProof(context.Background(), txHashes, &blockHash)
	if err != nil {
		return
	}
	fmt.Printf("tx proof indices: %v\n", txProof.Proof.Indices)
	fmt.Printf("tx proof lemmas size: %v", len(txProof.Proof.Lemmas)*32)
}

func filterBlock(client rpc.Client) {
	for index := fromBlock; index <= toBlock; index += 1 {
		block, err := client.GetBlockByNumber(context.Background(), index)
		if err != nil {
			break
		}
		if len(block.Transactions) < 50 {
			continue
		}
		fmt.Print("filter block number", block.Header.Number)
		fmt.Print("block hash", block.Header.Hash)
	}
}
