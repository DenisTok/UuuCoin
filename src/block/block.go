package block

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// 224 [index, timestamp, data, previousHash, hash]
type Block struct {
	index        uint64    // 64
	timestamp    time.Time // 64
	data         *data     // 32
	hash         [32]byte  // 32
	previousHash [32]byte  // 32
}

func NewBlock(previousBlock *Block, newBlockData *data) *Block {
	index := previousBlock.Index() + 1
	timestamp := time.Now().UTC()

	return &Block{
		index:        index,
		timestamp:    timestamp,
		data:         newBlockData,
		hash:         sha256.Sum256(calculateHash(previousBlock.hash, newBlockData.Hash(), index, timestamp)),
		previousHash: previousBlock.hash,
	}
}

func calculateHash(previousBlockHash [32]byte, newBlockDataHash [32]byte, index uint64, timestamp time.Time) []byte {
	t := timestamp.UnixNano()

	dataBytes := newBlockDataHash[:]

	previousBlockBytes := previousBlockHash[:]

	enc := []byte{
		byte(index >> 56), // bytes 0-7: index
		byte(index >> 48),
		byte(index >> 40),
		byte(index >> 32),
		byte(index >> 24),
		byte(index >> 16),
		byte(index >> 8),
		byte(index),
		byte(t >> 56), // bytes 8-15: timestamp
		byte(t >> 48),
		byte(t >> 40),
		byte(t >> 32),
		byte(t >> 24),
		byte(t >> 16),
		byte(t >> 8),
		byte(t),
	}

	enc = append(enc, dataBytes...)
	enc = append(enc, previousBlockBytes...)

	return enc
}

func (b *Block) Index() uint64 {
	return b.index
}

func (b *Block) Data() *data {
	return b.data
}

func (b *Block) Timestamp() time.Time {
	return b.timestamp
}

func (b *Block) Hash() [32]byte {
	return b.hash
}

func (b *Block) PreviousHash() [32]byte {
	return b.previousHash
}

func (b *Block) Print() {
	fmt.Println("** Block **")
	fmt.Printf("Index:\t%d\n", b.Index())
	fmt.Printf("Created:\t%s (unix %d)\n",
		b.Timestamp().Format(time.RFC3339Nano),
		b.Timestamp().UnixNano(),
	)
	if b.data != nil {
		fmt.Printf("With data:\t%s\n", b.Data().Info())
		fmt.Printf("Data hash:\t%x\n", b.Data().Hash())
	} else {
		fmt.Printf("With data:\t%s\n", "")
		fmt.Printf("Data hash:\t%x\n", sha256.Sum256([]byte("")))
	}
	fmt.Printf("Hash:\t%x\n", b.Hash())
	fmt.Printf("Previous Hash:\t%x\n", b.PreviousHash())
}
