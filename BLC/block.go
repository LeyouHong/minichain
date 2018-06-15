package BLC

import (
  "time"
)

type Block struct {
  // 1.height
  Height int64
  // 2.previous block Hash
  PrevBlockHash []byte
  // 3.transaction data
  Data []byte
  // 4.timestamp
  Timestamp int64
  // 5.currentblock Hash
  Hash []byte
  //nonce
  Nonce int64
}

// Set Hash
/*
func (block *Block) SetHash() {
  timeStr := strconv.FormatInt(block.Timestamp, 2)
  timeBytes := []byte(timeStr);

  heightBytes := utils.IntToHex(block.Height)

  blockBytes := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block.Data, timeBytes, block.Hash}, []byte{})
  hash := sha256.Sum256(blockBytes)
  block.Hash = hash[:]
}*/

// create new block
func NewBlock(data string, height int64, preBlockHash []byte) *Block {
  block := &Block{height, preBlockHash, []byte(data), time.Now().Unix(), nil, 0}
  //block.SetHash()
  return block
}

func CreateGenesisBlock(data string) *Block {
  firstBlockHash := []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
  return NewBlock(data, 1, firstBlockHash)
}
