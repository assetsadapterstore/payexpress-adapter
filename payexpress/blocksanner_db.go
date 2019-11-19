package payexpress

import (
	"fmt"
	"github.com/blocktree/openwallet/openwallet"
)

//SaveLocalBlockHead 记录区块高度和hash到本地
func (bs *PESSBlockScanner) SaveLocalBlockHead(blockHeight uint64, blockHash string) error {

	if bs.BlockchainDAI == nil {
		return fmt.Errorf("Blockchain DAI is not setup ")
	}

	header := &openwallet.BlockHeader{
		Hash:   blockHash,
		Height: blockHeight,
		Fork:   false,
		Symbol: bs.wm.Symbol(),
	}

	return bs.BlockchainDAI.SaveCurrentBlockHead(header)
}

//GetLocalBlockHead 获取本地记录的区块高度和hash
func (bs *PESSBlockScanner) GetLocalBlockHead() (uint64, string, error) {

	if bs.BlockchainDAI == nil {
		return 0, "", fmt.Errorf("Blockchain DAI is not setup ")
	}

	header, err := bs.BlockchainDAI.GetCurrentBlockHead(bs.wm.Symbol())
	if err != nil {
		return 0, "", err
	}

	return header.Height, header.Hash, nil
}

//SaveLocalBlock 记录本地新区块
func (bs *PESSBlockScanner) SaveLocalBlock(blockHeader *Block) error {

	if bs.BlockchainDAI == nil {
		return fmt.Errorf("Blockchain DAI is not setup ")
	}

	header := &openwallet.BlockHeader{
		Hash:              blockHeader.Hash,
		Merkleroot:        "",
		Previousblockhash: blockHeader.PrevBlockHash,
		Height:            blockHeader.Height,
		Time:              uint64(blockHeader.Time),
		Symbol:            bs.wm.Symbol(),
	}

	return bs.BlockchainDAI.SaveLocalBlockHead(header)
}

//GetLocalBlock 获取本地区块数据
func (bs *PESSBlockScanner) GetLocalBlock(height uint64) (*Block, error) {

	if bs.BlockchainDAI == nil {
		return nil, fmt.Errorf("Blockchain DAI is not setup ")
	}

	header, err := bs.BlockchainDAI.GetLocalBlockHeadByHeight(height, bs.wm.Symbol())
	if err != nil {
		return nil, err
	}

	block := &Block{
		Hash:   header.Hash,
		Height: header.Height,
	}

	return block, nil
}

//SaveUnscanRecord 保存交易记录到钱包数据库
func (bs *PESSBlockScanner) SaveUnscanRecord(record *openwallet.UnscanRecord) error {

	if bs.BlockchainDAI == nil {
		return fmt.Errorf("Blockchain DAI is not setup ")
	}

	return bs.BlockchainDAI.SaveUnscanRecord(record)
}

//DeleteUnscanRecord 删除指定高度的未扫记录
func (bs *PESSBlockScanner) DeleteUnscanRecord(height uint64) error {

	if bs.BlockchainDAI == nil {
		return fmt.Errorf("Blockchain DAI is not setup ")
	}

	return bs.BlockchainDAI.DeleteUnscanRecordByHeight(height, bs.wm.Symbol())
}

func (bs *PESSBlockScanner) GetUnscanRecords() ([]*openwallet.UnscanRecord, error) {

	if bs.BlockchainDAI == nil {
		return nil, fmt.Errorf("Blockchain DAI is not setup ")
	}

	return bs.BlockchainDAI.GetUnscanRecords(bs.wm.Symbol())
}