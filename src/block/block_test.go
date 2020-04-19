package block

import (
	"testing"
	"time"
)

func BenchmarkNewBlock(b *testing.B) {
	b.ReportAllocs()
	f := NewBlock(&Block{}, NewData([]byte("Hello, world")))
	for i := 0; i < b.N; i++ {
		NewBlock(f, NewData([]byte("Hello, world")))
	}
}

func BenchmarkCalculations(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		calculateHash(&Block{}, &data{
			info: []byte("Hello, world"),
			hash: [32]byte{0xe6, 0x8f, 0xe7, 0x8e, 0x06, 0x47, 0x00, 0xfe, 0x6b, 0x98, 0xe4, 0x7d, 0xc0, 0x75, 0x8a, 0x4f, 0x96, 0x6b, 0xd0, 0x27, 0x29, 0x9b, 0x68, 0x56, 0x42, 0xc6, 0x07, 0xea, 0x37, 0x6b, 0x7d, 0x47},
		}, 1, &timestamp{
			created: time.Now(),
		})
	}
}
