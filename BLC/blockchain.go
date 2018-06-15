package BLC

type Blockchain struct {
	Blocks []*Block
}

func (blc *Blockchain) AddBlockToBlockchain(data string, height int64, preHash []byte) {
	newBlock := NewBlock(data, height, preHash)
	blc.Blocks = append(blc.Blocks, newBlock)
}

func CreatBlockchainWithGenesisBlock() *Blockchain {
	genesisBlock := CreateGenesisBlock("Genesis block ...")
	//fmt.Println(genesisBlock)
	return &Blockchain{[]*Block{genesisBlock}}
}