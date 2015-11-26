package blkparser

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestParseBlock(t *testing.T) {
	// Bitcoin genesis block.
	genesisBlock, err := hex.DecodeString("0100000000000000000000000000000000000000000000000000000000000000000000003ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a29ab5f49ffff001d1dac2b7c0101000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4d04ffff001d0104455468652054696d65732030332f4a616e2f32303039204368616e63656c6c6f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73ffffffff0100f2052a01000000434104678afdb0fe5548271967f1a67130b7105cd6a828e03909a67962e0ea1f61deb649f6bc3f4cef38c4f35504e51ec112de5c384df7ba0b8d578a4c702b6bf11d5fac00000000")
	if err != nil {
		t.Error(err)
	}

	block, err := NewBlock(genesisBlock)
	if err != nil {
		t.Error(err)
	}

	if bytes.Equal(genesisBlock, block.Raw) != true {
		t.Error("For raw block",
			"expected", genesisBlock, "got", block.Raw)
	}
	if block.Hash != "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f" {
		t.Error("For block hash, expected 000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f, got", block.Hash)
	}
	if block.Version != 1 {
		t.Error("For block version, expected 1, got", block.Version)
	}
	if block.MerkleRoot != "3ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a" {
		t.Error("For merkle root, expected 3ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a, got", block.MerkleRoot)
	}
	if block.BlockTime != 1231006505 {
		t.Error("For block time, expected 1231006505, got", block.BlockTime)
	}
	if block.Bits != 0x1d00ffff {
		t.Error("For block bits, expected 0x1d00ffff, got", block.Bits)
	}
	if block.Nonce != 2083236893 {
		t.Error("For block nonce, expected 2083236893, got", block.Nonce)
	}
	if block.Size != 285 {
		t.Error("For block size, expected 285, got", block.Size)
	}
	if block.TxCnt != 1 {
		t.Error("For block tx count, expected 1, got", block.TxCnt)
	}
}
