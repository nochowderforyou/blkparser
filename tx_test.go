package blkparser

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestParseTx(t *testing.T) {
	// Bitcoin genesis block transaction.
	rawTx, err := hex.DecodeString("01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4d04ffff001d0104455468652054696d65732030332f4a616e2f32303039204368616e63656c6c6f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73ffffffff0100f2052a01000000434104678afdb0fe5548271967f1a67130b7105cd6a828e03909a67962e0ea1f61deb649f6bc3f4cef38c4f35504e51ec112de5c384df7ba0b8d578a4c702b6bf11d5fac00000000")
	if err != nil {
		t.Error(err)
	}
	tx, _ := NewTx(rawTx)

	if tx.Version != 1 {
		t.Error("For Tx version, expected 1, got", tx.Version)
	}

	// Test Tx input.
	if tx.TxInCnt != 1 {
		t.Error("For TxIn count, expected 1, got", tx.TxInCnt)
	}

	txIn := tx.TxIns[0]
	if txIn.InputHash != "0000000000000000000000000000000000000000000000000000000000000000" {
		t.Error("For tx input hash, expected 0000000000000000000000000000000000000000000000000000000000000000, got", txIn.InputHash)
	}
	if txIn.InputVout != 0xffffffff {
		t.Error("For tx input index, expected 0xffffffff, got", txIn.InputVout)
	}
	actualScriptSig, _ := hex.DecodeString("04ffff001d0104455468652054696d65732030332f4a616e2f32303039204368616e63656c6c6f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73")
	if bytes.Equal(txIn.ScriptSig, actualScriptSig) != true {
		t.Error("For tx input script, expected 04ffff001d0104455468652054696d65732030332f4a616e2f32303039204368616e63656c6c6f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73, got", txIn.ScriptSig)
	}
	if txIn.Sequence != 0xffffffff {
		t.Error("For tx input sequence, expected 0xffffffff, got", txIn.Sequence)
	}

	// Test Tx output.
	if tx.TxOutCnt != 1 {
		t.Error("For TxOut count, expected 1, got", tx.TxOutCnt)
	}

	txOut := tx.TxOuts[0]
	if txOut.Value != 5000000000 {
		t.Error("For tx output value, expected 5000000000, got", txOut.Value)
	}
	actualOutputScript, _ := hex.DecodeString("4104678afdb0fe5548271967f1a67130b7105cd6a828e03909a67962e0ea1f61deb649f6bc3f4cef38c4f35504e51ec112de5c384df7ba0b8d578a4c702b6bf11d5fac")
	if bytes.Equal(actualOutputScript, txOut.Pkscript) != true {
		t.Error("For tx output script, expected 4104678afdb0fe5548271967f1a67130b7105cd6a828e03909a67962e0ea1f61deb649f6bc3f4cef38c4f35504e51ec112de5c384df7ba0b8d578a4c702b6bf11d5fac, got", txOut.Pkscript)
	}

	if tx.LockTime != 0 {
		t.Error("For Tx locktime, expected 0, got", tx.LockTime)
	}

}
