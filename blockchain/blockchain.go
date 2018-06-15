package blockchain

import (
	"github.com/minichain/BLC"
	"github.com/minichain/consensus"
)

type Blockchain struct {
	Blocks []*BLC.Block
}

func (blc *Blockchain) AddBlockToBlockchain(data string, height int64, preHash []byte) {
	newBlock := BLC.NewBlock(data, height, preHash)
	//proof of work and return valid hash and nonce
	pow := consensus.NewProofOfWork(newBlock)
	hash, nonce := pow.Run()
	newBlock.Hash = hash[:]
	newBlock.Nonce = nonce
	blc.Blocks = append(blc.Blocks, newBlock)
}

func CreatBlockchainWithGenesisBlock() *Blockchain {
	genesisBlock := BLC.CreateGenesisBlock("Genesis block ...")
	//fmt.Println(genesisBlock)
	return &Blockchain{[]*BLC.Block{genesisBlock}}
}