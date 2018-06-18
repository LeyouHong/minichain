package BLC

import (
	"testing"
)

func TestSSerializeAndDeSerialize(t *testing.T) {
	firstBlockHash := []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	block := NewBlock("Serialize and DeSerialize", 1, firstBlockHash)
	blockbytes := block.Serialize()
	//fmt.Println(blockbytes)
	block = DeSerializeBlock(blockbytes)
	//fmt.Println(block)
}
