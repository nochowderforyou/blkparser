package blkparser

import (
	"encoding/binary"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"strings"
)

// ClamMainNet Indicates the CLAM network.
const ClamMainNet wire.BitcoinNet = 0x15352203

// ClamParams contains CLAM network parameters.
var ClamParams = chaincfg.Params{
	Name:             "clammainnet",
	Net:              ClamMainNet,
	DefaultPort:      "31174",
	PubKeyHashAddrID: 0x89,
	ScriptHashAddrID: 0x0d,
	PrivateKeyID:     0x85,
}

// Tx models the data in a transaction.
type Tx struct {
	Hash     string
	Size     uint32
	LockTime uint32
	Version  uint32
	Time     uint32
	Comment  string
	TxInCnt  uint32
	TxOutCnt uint32
	TxIns    []*TxIn
	TxOuts   []*TxOut
}

// TxIn models the data in a transaction input.
type TxIn struct {
	InputHash string
	InputVout uint32
	ScriptSig []byte
	Sequence  uint32
}

// TxOut models the data in a transaction output.
type TxOut struct {
	Addr     string
	Value    uint64
	Pkscript []byte
}

// ParseTxs deserializes a byte slice into a Tx slice.
func ParseTxs(txsraw []byte) (txs []*Tx, err error) {
	offset := int(0)
	txcnt, txcnt_size := DecodeVariableLengthInteger(txsraw[offset:])
	offset += txcnt_size

	txs = make([]*Tx, txcnt)

	txoffset := int(0)
	for i := range txs {
		txs[i], txoffset = NewTx(txsraw[offset:])
		txs[i].Hash = GetShaString(txsraw[offset : offset+txoffset])
		txs[i].Size = uint32(txoffset)
		offset += txoffset
	}

	return
}

// NewTx deserializes a byte slice into a Tx.
func NewTx(rawtx []byte) (tx *Tx, offset int) {
	tx = new(Tx)
	tx.Version = binary.LittleEndian.Uint32(rawtx[0:4])
	tx.Time = binary.LittleEndian.Uint32(rawtx[4:8])
	offset = 8

	txincnt, txincntsize := DecodeVariableLengthInteger(rawtx[offset:])
	offset += txincntsize

	tx.TxInCnt = uint32(txincnt)
	tx.TxIns = make([]*TxIn, txincnt)

	txoffset := int(0)
	for i := range tx.TxIns {
		tx.TxIns[i], txoffset = NewTxIn(rawtx[offset:])
		offset += txoffset

	}

	txoutcnt, txoutcntsize := DecodeVariableLengthInteger(rawtx[offset:])
	offset += txoutcntsize

	tx.TxOutCnt = uint32(txoutcnt)
	tx.TxOuts = make([]*TxOut, txoutcnt)

	for i := range tx.TxOuts {
		tx.TxOuts[i], txoffset = NewTxOut(rawtx[offset:])
		offset += txoffset
	}

	tx.LockTime = binary.LittleEndian.Uint32(rawtx[offset : offset+4])
	offset += 4
	comment, commentsize := DecodeVariableLengthInteger(rawtx[offset:])
	offset += commentsize
	commentstr := fmt.Sprintf("%+q", rawtx[offset:offset+comment])
	// Trim once on each end of the comment in case the actual speech has quotation marks.
	commentstr = strings.TrimSuffix(strings.TrimPrefix(commentstr, "\""), "\"")
	tx.Comment = commentstr
	offset += comment

	return
}

// NewTxIn deserializes a byte slice into a TxIn.
func NewTxIn(txinraw []byte) (txin *TxIn, offset int) {
	txin = new(TxIn)
	txin.InputHash = HashString(txinraw[0:32])
	txin.InputVout = binary.LittleEndian.Uint32(txinraw[32:36])
	offset = 36

	scriptsig, scriptsigsize := DecodeVariableLengthInteger(txinraw[offset:])
	offset += scriptsigsize

	txin.ScriptSig = txinraw[offset : offset+scriptsig]
	offset += scriptsig

	txin.Sequence = binary.LittleEndian.Uint32(txinraw[offset : offset+4])
	offset += 4
	return
}

// NewTxOut deserializes a byte slice into a TxOut.
func NewTxOut(txoutraw []byte) (txout *TxOut, offset int) {
	txout = new(TxOut)
	txout.Value = binary.LittleEndian.Uint64(txoutraw[0:8])
	offset = 8

	pkscript, pkscriptsize := DecodeVariableLengthInteger(txoutraw[offset:])
	offset += pkscriptsize

	txout.Pkscript = txoutraw[offset : offset+pkscript]
	offset += pkscript

	_, addrhash, _, err := txscript.ExtractPkScriptAddrs(txout.Pkscript, &ClamParams)
	if err != nil {
		return
	}
	if len(addrhash) != 0 {
		txout.Addr = addrhash[0].EncodeAddress()
	} else {
		txout.Addr = ""
	}

	return
}
