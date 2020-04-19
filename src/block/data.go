package block

import "crypto/sha256"

type data struct {
	info []byte
	hash [32]byte
}

func NewData(info []byte) *data {
	return &data{
		info: info,
		hash: sha256.Sum256(info),
	}
}

func (d *data) Hash() [32]byte {
	return d.hash
}

func (d *data) Info() []byte {
	if d.info == nil {
		return make([]byte, 0)
	}
	return d.info
}
