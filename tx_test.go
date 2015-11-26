package blkparser

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestParseTx(t *testing.T) {
	// Bitcoin genesis block transaction.
	// Arbitrarily chosen CLAM transaction.
	rawTx, err := hex.DecodeString("02000000f718575602d81735df4e8a2c7058f7a621d2779861a6df9874bd4adbfe68630db33951f561000000006a47304402204eb9cde1ed2058aae88048ca5ade9de723fa3130b409b1d35fa42f3db3f71e420220395ed7cdbcff6639130d7c4e99c9cd1ae85978213a24a0b12e9b4a0f20bed8540121020639bc1d4a121d47668e0a620bc3982fa0df05e0cb311148e07ec7e03d9ffa4affffffff5945170f033aa2799a5363fc337a813d1721156e5a5889dfcf83be7b85364ee12300000049483045022100a2ac412de128d128c184395e09a926ca6a8a0b58e7c835d83727655a8ec1d22002204f8123c9121fc04a6d543c9e2f3b87e4765ca0770836c816aec45761977968ea01ffffffff0277890204000000001976a9146ca9d6964363832f8e1488a4d686625f300ed49b88ac20069836010000001976a9140b8364fa470f2f3c9b359e72c8034ec957e46ea488ac0000000011676f6f646c75636b2070656173656e697a")
	if err != nil {
		t.Error(err)
	}
	tx, _ := NewTx(rawTx)

	if tx.Version != 2 {
		t.Error("For Tx version, expected 2, got", tx.Version)
	}
	if tx.Time != 1448548599 {
		t.Error("For Tx time, expected 1448548599, got", tx.Time)
	}

	// Test Tx input.
	if tx.TxInCnt != 2 {
		t.Error("For TxIn count, expected 2, got", tx.TxInCnt)
	}

	txIn := tx.TxIns[0]
	if txIn.InputHash != "61f55139b30d6368fedb4abd7498dfa6619877d221a6f758702c8a4edf3517d8" {
		t.Error("For tx input 0 hash, expected 61f55139b30d6368fedb4abd7498dfa6619877d221a6f758702c8a4edf3517d8, got", txIn.InputHash)
	}
	if txIn.InputVout != 0 {
		t.Error("For tx input 0 index, expected 0, got", txIn.InputVout)
	}
	actualScriptSig, _ := hex.DecodeString("47304402204eb9cde1ed2058aae88048ca5ade9de723fa3130b409b1d35fa42f3db3f71e420220395ed7cdbcff6639130d7c4e99c9cd1ae85978213a24a0b12e9b4a0f20bed8540121020639bc1d4a121d47668e0a620bc3982fa0df05e0cb311148e07ec7e03d9ffa4a")
	if bytes.Equal(txIn.ScriptSig, actualScriptSig) != true {
		t.Errorf("For tx input 0 script, expected 47304402204eb9cde1ed2058aae88048ca5ade9de723fa3130b409b1d35fa42f3db3f71e420220395ed7cdbcff6639130d7c4e99c9cd1ae85978213a24a0b12e9b4a0f20bed8540121020639bc1d4a121d47668e0a620bc3982fa0df05e0cb311148e07ec7e03d9ffa4a, got %x", txIn.ScriptSig)
	}
	if txIn.Sequence != 0xffffffff {
		t.Error("For tx input 0 sequence, expected 0xffffffff, got", txIn.Sequence)
	}

	// Test Tx output.
	if tx.TxOutCnt != 2 {
		t.Error("For TxOut count, expected 2, got", tx.TxOutCnt)
	}

	txOut := tx.TxOuts[0]
	if txOut.Addr != "xJDCLAMZW9xZEFRfWhxXkcCGYkHAzhsjT5" {
		t.Error("For tx output 0 address, expected xJDCLAMZW9xZEFRfWhxXkcCGYkHAzhsjT5, got", txOut.Addr)
	}
	if txOut.Value != 67275127 {
		t.Error("For tx output 0 value, expected 67275127, got", txOut.Value)
	}
	actualOutputScript, _ := hex.DecodeString("76a9146ca9d6964363832f8e1488a4d686625f300ed49b88ac")
	if bytes.Equal(actualOutputScript, txOut.Pkscript) != true {
		t.Errorf("For tx output 0 script, expected 76a9146ca9d6964363832f8e1488a4d686625f300ed49b88ac, got %x", txOut.Pkscript)
	}

	if tx.LockTime != 0 {
		t.Error("For Tx locktime, expected 0, got", tx.LockTime)
	}

	actualComment := "goodluck peaseniz"
	if actualComment != tx.Comment {
		t.Error("For tx comment, expected", actualComment, "got", tx.Comment)
	}
}
