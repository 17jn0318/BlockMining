package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

type Block struct {
	Version       int64
	PrevBlockHash []byte
	Hash          []byte
	TimeStamp     int64
	Difficulty    int64
	Nonce         int64 //ここまで、 Heads
	Data          []byte
}

type Blockchain struct {
	blocks []*Block
}

func NewBlockchain() *Blockchain { // (1)Create a Blockchain
	return &Blockchain{[]*Block{NewGenesisBlock()}} //add block into blockchain
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlock(data string, prevBlockHash []byte) *Block { //create block
	block := Block{Version: 1,
		PrevBlockHash: prevBlockHash,
		TimeStamp:     time.Now().Unix(),
		Difficulty:    difficulty,
		Nonce:         0,
		Data:          []byte(data)}
	//block.SetHash()
	pow := NewProofOfWork(&block)

	nonce, hash := pow.Min()
	block.Nonce = nonce
	block.Hash = hash
	return &block
}
func IntToByte(num int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	CheckErr("IntToByte", err)
	return buffer.Bytes()
}

func CheckErr(pos string, err error) {
	if err != nil {
		fmt.Println("pos,err", pos, err)
		os.Exit(1)
	}
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1] //first block's index is 0
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock) //add block into blockchain
}

func main() {
	bc := NewBlockchain() // Create a blockchain

	bc.AddBlock("Masato Fukaya")
	bc.AddBlock("Joao Amaral")

	for _, block := range bc.blocks {
		fmt.Printf("Version: %d\n", block.Version)
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("TimeStamp : %d\n", block.TimeStamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Bits: %d\n", block.Difficulty)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Println()
	}
}
