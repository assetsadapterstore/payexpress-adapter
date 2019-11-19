package payexpress

import (
	"github.com/blocktree/go-owcrypt"
	"github.com/blocktree/openwallet/common/file"
	"path/filepath"
	"strings"
)

const (
	//币种
	Symbol    = "PESS"
	CurveType = owcrypt.ECC_CURVE_ED25519
	PESS_SEQUENCEID_KEY = "PESS_SEQUENCEID_KEY"

	//默认配置内容
	defaultConfig = `

# RPC api url
serverAPI = ""
`
)

type WalletConfig struct {

	//币种
	Symbol string
	//配置文件路径
	configFilePath string
	//配置文件名
	configFileName string
	//区块链数据文件
	//BlockchainFile string
	//本地数据库文件路径
	dbPath string
	//钱包服务API
	ServerAPI string
	//默认配置内容
	DefaultConfig string
	//曲线类型
	CurveType uint32
	//链ID
	NetworkID string
	//固定手续费
	FixFees string
	//激活账户最低余额
	ActiveBalance string
	//数据目录
	DataDir string
}

func NewConfig(symbol string) *WalletConfig {

	c := WalletConfig{}

	//币种
	c.Symbol = symbol
	c.CurveType = CurveType

	//区块链数据
	//blockchainDir = filepath.Join("data", strings.ToLower(Symbol), "blockchain")
	//配置文件路径
	c.configFilePath = filepath.Join("conf")
	//配置文件名
	c.configFileName = c.Symbol + ".ini"
	//区块链数据文件
	//c.BlockchainFile = "blockchain.db"
	//本地数据库文件路径
	c.dbPath = filepath.Join("data", strings.ToLower(c.Symbol), "db")
	//钱包服务API
	c.ServerAPI = ""
	//固定手续费
	c.FixFees = "0"
	//激活账户最低余额
	c.ActiveBalance = "0"

	//创建目录
	//file.MkdirAll(c.dbPath)

	return &c
}

//创建文件夹
func (wc *WalletConfig) makeDataDir() {

	if len(wc.DataDir) == 0 {
		//默认路径当前文件夹./data
		wc.DataDir = "data"
	}

	//本地数据库文件路径
	wc.dbPath = filepath.Join(wc.DataDir, strings.ToLower(wc.Symbol), "db")

	//创建目录
	file.MkdirAll(wc.dbPath)
}