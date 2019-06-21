# payexpress-adapter

payexpress-adapter适配了openwallet.AssetsAdapter接口，给应用提供了底层的区块链协议支持。

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建PESS.ini文件，编辑如下内容：

```ini

# RPC api url
;serverAPI = "http://mainnet.payexpress.io/api/v1/"
serverAPI = "https://testnet-sebak.blockchainos.org/api/v1/"
# fix fees for transaction
fixFees = "0.001"
# To make account valid, should deposit minimum balance, default = 0.1
activeBalance = "0.1"
# PESS networkID, default(mainnet) networkID = "sebak-payexpress-network; 2018-12-18",
;networkID = "sebak-payexpress-network; 2018-12-18"
networkID = "sebak-test-network"
# Cache data file directory, default = "", current directory: ./data
dataDir = ""

```

## 资料介绍

### 官网

http://www.payexpress.io/

### 区块浏览器

https://explorer.payexpress.io/
http://explorer-pess.s3-website.ap-northeast-2.amazonaws.com/

### RPC API

#### mainnet

http://mainnet.payexpress.io/

#### testnet

https://testnet-sebak.blockchainos.org/

### github

https://github.com/bosnet/sebak

### 适配资料

#### rpc api文档的链接

http://devteam.blockchainos.org/docs/api
