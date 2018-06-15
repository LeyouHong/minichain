package BLC

import (
	"testing"
)

func TestNewBlock(t *testing.T) {
	blockchain := CreatBlockchainWithGenesisBlock()
	blockchain.AddBlockToBlockchain("send 100eth to Leyou",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("send 1000eth to Leyou",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("send 10000eth to Leyou",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash)

	//for i:=0; i < len(blockchain.Blocks); i++ {
	//	fmt.Println(blockchain.Blocks[i])
	//}
}
