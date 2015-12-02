package blkparser

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestParseBlock(t *testing.T) {
	// Arbitrarily chosen CLAM block.
	rawBlock, err := hex.DecodeString("070000002a120f263757acb9d30ef8d03e0087f54a47e498aa4911edda2d45622f9ed8aec65c5b872b572dbfae9eab50bb2a74263803a51a7748e99e48f7645df50c6b55c020575676b6001b000000000302000000c0205756010000000000000000000000000000000000000000000000000000000000000000ffffffff04036c650bffffffff01000000000000000000000000000002000000c0205756019dcb2ad4cb7af4b737f2e3764c8aab58629eae9502b015343c8dec0165888e1101000000494830450221009d2f091d0cc3a06ac610e2a9c998ff37b8decc445af6355d7ad3e8ce3900ff5502203e2adeeda8befab3cb3747021287d5c964b388a64e34fa2e96d2e4c1081c897001ffffffff0200000000000000000020ed4d8b000000002321037bedfabb451755cf6061636c8004dba32cb95095ba8cba61de236a70f95e3d2aac000000003445787072657373696f6e206f6620506f6c69746963616c2046726565646f6d3a20416e6172636869737420636f6d6d756e69736d020000008f205756014e000e0a6cabefa4dbfebb5464780e888d005b3b493ea18e9d98c6c3f8dbdae5010000006b483045022100d65bd7f42fc4afa41a20286b6cf7daec80a4f86f6f2d67f70283b2b32daca98102206250bfec87231f891cf40830af9599879b9ce19a20bf589a73288b6468613e640121024051f34c45a7501169ba6f81eafd0bcaa2d3def9988071610c0410cca4fd51b5ffffffff06c39c0000000000001976a914fe2ed72632623d2f9a32c440799831c84c1db96b88aca60e0100000000001976a9147ae7e5015a5de27dc08b1ff53d2de040808a122b88acb7ec0000000000001976a914c495dbcd771510d52cc2eeaff5d093cae570703888ac05d90000000000001976a914bf460f324dd8a6652703f053cfff65927f6d650088acadf40000000000001976a914125f90e475d54d17a4556fbf2c65d4db65e61eae88acc0000835000000001976a91494c533fe25e6211b75303ac6fe253ffe4b248ae388ac000000002645787072657373696f6e206f662052656c6967696f75732046726565646f6d3a204b756c616d46304402203241198edb3ced65052cfd4c1a6976bc2992c42d3f11d351e06e1c479cce1c3c022072f791601210b489bc496415c4c3d06dd9c8de4479f16b7f4b9581e91906d661")
	if err != nil {
		t.Error(err)
	}

	block, err := NewBlock(rawBlock)
	if err != nil {
		t.Error(err)
	}

	if bytes.Equal(rawBlock, block.Raw) != true {
		t.Errorf("For raw block, expected %x, got %x", rawBlock, block.Raw)
	}
	if block.Hash != "7727fa0415d44591de16cfa4727f95f7af027c8d8c99acf345171bff89bac8b4" {
		t.Error("For block hash, expected 7727fa0415d44591de16cfa4727f95f7af027c8d8c99acf345171bff89bac8b4, got", block.Hash)
	}
	if block.Version != 7 {
		t.Error("For block version, expected 7, got", block.Version)
	}
	if block.MerkleRoot != "556b0cf55d64f7489ee948771aa5033826742abb50ab9eaebf2d572b875b5cc6" {
		t.Error("For merkle root, expected 556b0cf55d64f7489ee948771aa5033826742abb50ab9eaebf2d572b875b5cc6, got", block.MerkleRoot)
	}
	if block.BlockTime != 1448550592 {
		t.Error("For block time, expected 1448550592, got", block.BlockTime)
	}
	if block.Bits != 0x1b00b676 {
		t.Error("For block bits, expected 0x1b00b676, got", block.Bits)
	}
	if block.Nonce != 0 {
		t.Error("For block nonce, expected 0, got", block.Nonce)
	}
	if block.Size != 860 {
		t.Error("For block size, expected 860, got", block.Size)
	}
	if block.TxCnt != 3 {
		t.Error("For block tx count, expected 3, got", block.TxCnt)
	}
	actualSig := "304402203241198edb3ced65052cfd4c1a6976bc2992c42d3f11d351e06e1c479cce1c3c022072f791601210b489bc496415c4c3d06dd9c8de4479f16b7f4b9581e91906d661"
	if block.BlockSig != actualSig {
		t.Error("For block sig, expected", actualSig, "got", block.BlockSig)
	}
	if block.IsProofOfStake() != true {
		t.Error("For block IsProofOfStake(), expected true, got false")
	}
}
