package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Timestamp     string //timeNow
	Data          []byte //blockData
	PrevBlockHash []byte //prevs hash
	Hash          []byte // hash of this
}

type Blockchain struct {
	blocks []*Block
}

func NewBlockchain() *Blockchain { // (1)Create a Blockchain
	return &Blockchain{[]*Block{NewDefaultBlock()}} //add block into blockchain
}

func NewDefaultBlock() *Block { //(2)create a default block 最初のBlockビットコインでの始まりは”Genesis Block”っていうBlock
	return NewBlock("Jin xinzhe", []byte{})
}

func NewBlock(data string, prevBlockHash []byte) *Block { //create block
	block := &Block{time.Now().String(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

func (b *Block) SetHash() {
	timestamp := []byte(b.Timestamp) //[]byte type Cast

	//func Join(s [][]byte, sep []byte) []byte
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{}) //link and create new []byte slice
	hash := sha256.Sum256(headers)                                                //hashvalue
	b.Hash = hash[:]                                                              //hashslice
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
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
