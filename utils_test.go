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

func TestGetScryptString(t *testing.T) {
	// Litecoin block header from litecoin wiki.
	rawBlock, err := hex.DecodeString("01000000f615f7ce3b4fc6b8f61e8f89aedb1d0852507650533a9e3b10b9bbcc30639f279fcaa86746e1ef52d3edb3c4ad8259920d509bd073605c9bf1d59983752a6b06b817bb4ea78e011d012d59d4")
	if err != nil {
		t.Error(err)
	}

	actualHash := "0000000110c8357966576df46f3b802ca897deb7ad18b12f1c24ecff6386ebd9"
	blockHash, err := GetScryptString(rawBlock)
	if err != nil {
		t.Error(err)
	}
	if actualHash != blockHash {
		t.Error("For scrypt string of litecoin block",
			"expected", actualHash, "got", blockHash)
	}
}
