package payexpress

import (
	"fmt"
	"github.com/blocktree/openwallet/common"
	"github.com/blocktree/openwallet/crypto"
	"github.com/blocktree/openwallet/openwallet"
	"github.com/tidwall/gjson"
	"math/big"
	"time"
)

type AddrBalance struct {
	Address      string
	Balance      *big.Int
	TokenBalance *big.Int
	SequenceID   uint64
}

func NewAddrBalance(result *gjson.Result) *AddrBalance {
	obj := AddrBalance{}
	obj.Address = result.Get("address").String()
	obj.Balance = big.NewInt(result.Get("balance").Int())
	obj.SequenceID = result.Get("sequence_id").Uint()
	return &obj
}

type txFeeInfo struct {
	GasUsed  *big.Int
	GasPrice *big.Int
	Fee      *big.Int
}

func (f *txFeeInfo) CalcFee() error {
	fee := new(big.Int)
	fee.Mul(f.GasUsed, f.GasPrice)
	f.Fee = fee
	return nil
}

type Block struct {
	Hash             string
	Version          uint64
	PrevBlockHash    string
	TransactionsRoot string
	Time             int64
	Height           uint64
	tx               []string
}

func NewBlock(result *gjson.Result) *Block {
	obj := Block{}
	//解析json
	obj.Hash = result.Get("hash").String()
	obj.Version = result.Get("version").Uint()
	obj.PrevBlockHash = result.Get("prev_block_hash").String()
	obj.Height = result.Get("height").Uint()
	t, _ := time.Parse(time.RFC3339Nano, result.Get("proposed_time").String())
	obj.Time = t.Unix()

	txs := make([]string, 0)
	for _, tx := range result.Get("transactions").Array() {
		txs = append(txs, tx.String())
	}

	obj.tx = txs

	return &obj
}

//BlockHeader 区块链头
func (b *Block) BlockHeader(symbol string) *openwallet.BlockHeader {

	obj := openwallet.BlockHeader{}
	//解析json
	obj.Hash = b.Hash
	//obj.Confirmations = b.Confirmations
	//obj.Merkleroot = b.TransactionMerkleRoot
	obj.Previousblockhash = b.PrevBlockHash
	obj.Height = b.Height
	obj.Version = uint64(b.Version)
	obj.Time = uint64(obj.Time)
	obj.Symbol = symbol

	return &obj
}

type Transaction struct {
	BlockHeight uint64
	BlockHash   string
	Time        int64
	Fee         *big.Int
	TxID        string
	Source      string
	SequenceId  uint64
	Operations  []*Operation
}

func NewTransaction(result *gjson.Result) *Transaction {
	obj := Transaction{}
	obj.BlockHash = result.Get("block").String()
	obj.Fee = big.NewInt(result.Get("fee").Int())
	obj.TxID = result.Get("hash").String()
	obj.Source = result.Get("source").String()
	t, _ := time.Parse(time.RFC3339Nano, result.Get("created").String())
	obj.Time = t.Unix()

	operations := make([]*Operation, 0)
	if result.Get("operations").IsArray() {
		for _, o := range result.Get("operations").Array() {
			operations = append(operations, NewOperation(o, &obj))
		}
	}
	obj.Operations = operations

	return &obj
}

type Operation struct {
	BlockHeight uint64
	BlockHash   string
	Time        int64
	Fee         *big.Int
	TxID        string
	Source      string
	Type        string
	Target      string
	Amount      *big.Int
}

func NewOperation(result gjson.Result, tx *Transaction) *Operation {
	obj := Operation{}
	obj.Type = result.Get("H.type").String()
	obj.Target = result.Get("B.target").String()
	obj.Amount = big.NewInt(result.Get("B.amount").Int())
	obj.BlockHash = tx.BlockHash
	obj.Fee = tx.Fee
	obj.TxID = tx.TxID
	obj.Source = tx.Source
	obj.Time = tx.Time
	return &obj
}

//UnscanRecords 扫描失败的区块及交易
type UnscanRecord struct {
	ID          string `storm:"id"` // primary key
	BlockHeight uint64
	TxID        string
	Reason      string
}

func NewUnscanRecord(height uint64, txID, reason string) *UnscanRecord {
	obj := UnscanRecord{}
	obj.BlockHeight = height
	obj.TxID = txID
	obj.Reason = reason
	obj.ID = common.Bytes2Hex(crypto.SHA256([]byte(fmt.Sprintf("%d_%s", height, txID))))
	return &obj
}
