package main

import (
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockchain() // 제네시스 블록을 만듦과 동시에 새로운 블록체인을 만든다.

	bc.AddNewBlock("Send 1 BTC to Ivan")
	// "Send 1 BTC to Ivan" 라는 Data를 가지고, Genesisblock의 Hash값을 PrecBlockHash로 가지는 블록을 생성한다.

	bc.AddNewBlock("Send 2 more BTC to Ivan")
	// "Send 2 more BTC to Ivan" 라는 Data를 가지고, "Send 1 BTC to Ivan"의 Data를 가진 블록(1번째 블록)의 Hash값을 PrevBlockHash값으로 가지는 블록을 생성한다.

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

// git 관련 질문
// add를 커밋 할때마다 해야 하는가?
// branch 생성에 대해서 다시 설명해주셨으면..
