package blkparser

import (
	"encoding/hex"
	"testing"
)

func TestGetShaString(t *testing.T) {
	// Bitcoin genesis block header.
	rawBlock, err := hex.DecodeString("0100000000000000000000000000000000000000000000000000000000000000000000003ba3edfd7a7b12b27ac72c3e67768f617fc81bc3888a51323a9fb8aa4b1e5e4a29ab5f49ffff001d1dac2b7c")
	if err != nil {
		t.Error(err)
	}

	actualHash := "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"
	blockHash := GetShaString(rawBlock)
	if actualHash != blockHash {
		t.Error("For sha string of genesis block",
			"expected", actualHash, "got", blockHash)
	}
}
