package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}
type BlockChain struct {
	blocks []*Block
}

func main() {
	chain := InitBlockChain()
	chain.AddBlock("Vyas First Block")
	chain.AddBlock("Second Block ")
	chain.AddBlock("Third Block")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash : %x\n", block.PrevHash)
		fmt.Printf("Data %s\n", block.Data)
		fmt.Printf("Hash %x\n", block.Hash)

	}
}
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DervieHash()
	return block
}
func (b *Block) DervieHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)

}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
