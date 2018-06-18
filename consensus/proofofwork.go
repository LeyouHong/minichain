package consensus

import (
	"github.com/minichain/BLC"
	"math/big"
	"bytes"
	"github.com/minichain/utils"
	"crypto/sha256"
	"fmt"
)

const targetBit = 16

type ProofOfWork struct {
	Block *BLC.Block
	target *big.Int
}

func(proofOfWork *ProofOfWork) Run() ([]byte, int64) {
	nonce := int64(0)
	var hashInt big.Int // save new hash data
	var hash [32]byte

	for {
		dataBytes := proofOfWork.PrePareData(nonce)
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce = nonce + 1;
	}
	return hash[:], nonce
}

func (proofOfWork *ProofOfWork) IsValid() bool {
	var hashInt big.Int // save new hash data
	hashInt.SetBytes(proofOfWork.Block.Hash)

	if proofOfWork.target.Cmp(&hashInt) == 1 {
		return true
	}

	return false
}

func (proofOfWork *ProofOfWork) PrePareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			proofOfWork.Block.PrevBlockHash,
			proofOfWork.Block.Data,
			utils.IntToHex(proofOfWork.Block.Timestamp),
			utils.IntToHex(int64(targetBit)),
			utils.IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (proofOfWork *ProofOfWork) CreateNewPowBlock() *BLC.Block {
	hash, nonce := proofOfWork.Run()
	proofOfWork.Block.Hash = hash[:]
	proofOfWork.Block.Nonce = nonce

	return proofOfWork.Block
}

func NewProofOfWork (blc *BLC.Block) *ProofOfWork {
	target := big.NewInt(1)
	target = target.Lsh(target, 256 - targetBit)
	return &ProofOfWork{blc, target}
}