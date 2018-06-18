package BLC

import (
  "time"
  "bytes"
  "encoding/gob"
  "log"
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

func (block *Block) Serialize() []byte {
  var result bytes.Buffer

  encoder := gob.NewEncoder(&result)
  err := encoder.Encode(block)
  if err != nil {
    log.Panic(err)
  }

  return result.Bytes()
}

func DeSerializeBlock(blockBytes []byte) *Block {
  var block Block

  docoder := gob.NewDecoder(bytes.NewReader(blockBytes))
  err := docoder.Decode(&block)
  if err != nil {
    log.Panic(err)
  }

  return &block
}

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
