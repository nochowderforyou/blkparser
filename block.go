package blkparser

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

// BlockPos models the position of a block's data in blockchain files.
type BlockPos struct {
	FileId uint32 `json:"fileId"`
	Pos    int64  `json:"pos"`
}

// Block models the data in a blockchain block.
type Block struct {
	Raw         []byte   `json:"-"`
	Pos         BlockPos `json:"-"`
	Hash        string   `json:"hash"`
	Height      uint     `json:"height"`
	Txs         []*Tx    `json:"tx,omitempty"`
	Version     uint32   `json:"ver"`
	MerkleRoot  string   `json:"mrkl_root"`
	BlockTime   uint32   `json:"time"`
	Bits        uint32   `json:"bits"`
	Nonce       uint32   `json:"nonce"`
	Size        uint32   `json:"size"`
	BlockSig    string   `json:"signature"`
	TxCnt       uint32   `json:"n_tx"`
	TotalBTC    uint64   `json:"total_out"`
	BlockReward float64  `json:"-"`
	Parent      string   `json:"prev_block"`
	Next        string   `json:"next_block"`
}

func (block *Block) IsProofOfStake() bool {
	if len(block.Txs) > 1 && block.Txs[1].IsCoinStake() {
		return true
	}
	return false
}

// NewBlock deserializes a byte slice into a Block.
func NewBlock(rawblock []byte) (block *Block, err error) {
	block = new(Block)
	block.Raw = rawblock

	block.Version = binary.LittleEndian.Uint32(rawblock[0:4])
	if block.Version > 6 {
		block.Hash = GetShaString(rawblock[:80])
	} else {
		block.Hash, _ = GetScryptString(rawblock[:80])
	}
	if !bytes.Equal(rawblock[4:36], []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}) {
		block.Parent = HashString(rawblock[4:36])
	}
	block.MerkleRoot = HashString(rawblock[36:68])
	block.BlockTime = binary.LittleEndian.Uint32(rawblock[68:72])
	block.Bits = binary.LittleEndian.Uint32(rawblock[72:76])
	block.Nonce = binary.LittleEndian.Uint32(rawblock[76:80])
	block.Size = uint32(len(rawblock))

	txs, offset, _ := ParseTxs(rawblock[80:])

	block.Txs = txs
	block.TxCnt = uint32(len(txs))

	if block.IsProofOfStake() {
		offset += 80
		blocksig, blocksigsize := DecodeVariableLengthInteger(rawblock[offset:])
		offset += blocksigsize

		blocksigHex := hex.EncodeToString(rawblock[offset : offset+blocksig])
		block.BlockSig = blocksigHex
	}

	return
}
