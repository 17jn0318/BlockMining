package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

type ProofOfWork struct {
	block *Block

	target *big.Int
}

const difficulty = 24

func NewProofOfWork(block *Block) *ProofOfWork {

	target := big.NewInt(1)                  //0000000000000.....1
	target.Lsh(target, uint(256-difficulty)) //二進数で、２４桁は１６進数六桁
	pow := ProofOfWork{block: block, target: target}
	return &pow
}

func (pow *ProofOfWork) SetData(nonce int64) []byte {
	b := pow.block
	data := bytes.Join([][]byte{IntToByte(b.Version), b.PrevBlockHash, IntToByte(b.TimeStamp), IntToByte(difficulty), IntToByte(nonce), b.Data}, []byte{})
	return data
}

func (pow *ProofOfWork) Min() (int64, []byte) {
	var nonce int64 = 0
	var hash [32]byte
	var hashInt big.Int //target is big.int hasaはハッシュ値

	fmt.Println("begin mining.....")
	fmt.Printf("target :%x\n", pow.target.Bytes())

	for nonce < math.MaxInt64 {
		data := pow.SetData(nonce)
		hash = sha256.Sum256(data)

		hashInt.SetBytes(hash[:])

		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("found nonce, nonce : %d, hash : %x\n", nonce, hash)
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}
