# PrivatePOAEthereum

## 正式使用硬件需求

Hardware Requirements
Minimum:

CPU with 2+ cores
4GB RAM
1TB free storage space to sync the Mainnet
8 MBit/sec download Internet service
Recommended:

Fast CPU with 4+ cores
16GB+ RAM
High-performance SSD with at least 1TB of free space
25+ MBit/sec download Internet service

出块节点： 开放30303 （测试可以开放8545，后续需要关闭）
同步节点：开放30303 8545

## 安装golang gcc

```shell
sudo apt install build-essential
```

## Clone Ethereum Code

```shell
git clone https://github.com/ethereum/go-ethereum.git
```

##  编译geth工具

```shell
 cd go-ethereum && make geth
```

## 写入geth的环境变量
```shell
vi ~/.bashrc
```
```shell
export ETHPATH=/opt/go-ethereum/build/bin
export PATH=$ETHPATH:$PATH
```
```shell
source ~/.bashrc
```

## 生成账户地址 

私钥通过用户输入的密码加密存储,可以通过golang解析出账户地址和私钥： https://github.com/yuhuajing/PrivatePOAEthereum/blob/main/parsePrivateKey.go

```shell
geth account new --datadir ./node1
```
> 0x6593B47be3F4Bd1154c2faFb8Ad4aC4EFddD618f
> 797391c7bd2e156e52329ceb6471496798e0c125ef35c4c3393329bd2a64f3f5

```shell
geth account new --datadir ./node2
```
> 0x6C345f0771a2f2B2694f97522D3371bF87b6BDF9
> 08a7533871d3a2e01d3a8849320cbfb703eb20c5dd2a9ccd2d9780eba5659c8e
```shell
geth account new --datadir ./node3
```
> 0xab6bbb89eFd62dF605C881E692960a4951238D71
> b5cd29c38830904b433c52e5a8b8ac1a8349175afd8d6b8cf60333ad45ccdc1a

## 构造创世区块
```text
1. chainID:自定义的链ID
2. homesteadBlock、eip150Block、eip155Block、eip158Block、byzantiumBlock、constantinopleBlock、petersburgBlock：各项提案和升级的区块高度
3. period:出块时间间隔，0为不允许出空交易块，会等待有交易才出块
4. epoch:更新出块节点列表的周期
5. difficulty：POA下无作用
6. gasLimit：gasLimit限制
7. extradata：POA模式下用来指定验证者地址,账户地址去掉0x后加上64个前缀0和130个后缀0，比如0x+（64个前缀0）+5534F5024146D16a5C1ce60A9f5a2e9794e3F981+（130个后缀0）
8. alloc：用来预置账号以及账号的以太币数量，比如预置0x0B587FFD0BBa122fb5ddc19AD6eEcEB1D2dBbff7地址拥有1000ETH（1000*10^18WEI）
```
genesis.json
```json
{
   "config":{
      "chainId":12345,
      "homesteadBlock":0,
      "eip150Block":0,
      "eip155Block":0,
      "eip158Block":0,
      "byzantiumBlock":0,
      "constantinopleBlock":0,
      "petersburgBlock":0,
      "istanbulBlock":0,
      "berlinBlock":0,
      "londonBlock": 0,
      "clique":{
         "period":5,
         "epoch":300
      }
   },
   "alloc":{
      "0x6593B47be3F4Bd1154c2faFb8Ad4aC4EFddD618f":{
         "balance":"1000000000000000000000"
      },
      "0x6C345f0771a2f2B2694f97522D3371bF87b6BDF9":{
         "balance":"1000000000000000000000"
      },
      "0xab6bbb89eFd62dF605C881E692960a4951238D71":{
         "balance":"1000000000000000000000"
      }
   },
  "coinbase": "0x6593B47be3F4Bd1154c2faFb8Ad4aC4EFddD618f",
  "difficulty": "1",
  "extraData": "0x00000000000000000000000000000000000000000000000000000000000000006593B47be3F4Bd1154c2faFb8Ad4aC4EFddD618f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
  "gasLimit": "800000000",
  "nonce": "0x0000000000000000",
  "mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp": "0x00"
}
```
## 启动节点1

1. 创建创世区块
```shell
geth --datadir "/opt/etherData/node1" init /opt/etherData/genesis.json
```
2. 启动node1
node1为预置的矿工，因此在启动节点时，需要解锁节点并开启出块操作。
```shell
nohup geth --identity "myethereum" --datadir /opt/etherData/node1 --networkid 12345 --authrpc.port 8551 --http --http.port 8545  --http.corsdomain "*" --graphql --ws --ws.port 8546 --ws.addr 0.0.0.0 --ws.origins "*" --port 30303 --http.addr 0.0.0.0 --http.api "eth,net,web3,personal,admin,miner" --allow-insecure-unlock --rpc.enabledeprecatedpersonal --syncmode "full" --mine --miner.etherbase 0x6593b47be3f4bd1154c2fafb8ad4ac4efddd618f --unlock 0x6593b47be3f4bd1154c2fafb8ad4ac4efddd618f --keystore /opt/etherData/node1/keystore/ --password /opt/etherData/node1/password.txt --nodiscover >> geth1.log 2>&1 &
```

## 启动节点2

1. 创建创世区块
```shell
geth --datadir "/opt/etherData/node2" init /opt/etherData/genesis.json
```
2. 启动node2
```shell
nohup geth --datadir "/opt/etherData/node2" --networkid 12345 --authrpc.port 8552 --http --http.port 8547 --http.corsdomain "*" --ws --ws.port 8548 --ws.addr 0.0.0.0 --ws.origins "*" --port 30304 --http.addr 0.0.0.0 --http.api "eth,net,web3,personal,admin,miner" --allow-insecure-unlock --rpc.enabledeprecatedpersonal --syncmode "full" >> geth2.log 2>&1 &
```

## 启动节点3

1. 创建创世区块
```shell
geth --datadir "/opt/etherData/node3" init /opt/etherData/genesis.json
```
2. 启动node3
```shell
nohup geth --datadir "/opt/etherData/node3" --networkid 12345 --authrpc.port 8553 --http --http.port 8549 --http.corsdomain "*" --ws --ws.addr 0.0.0.0 --ws.port 8550 --port 30305 --ws.addr 0.0.0.0 --ws.origins "*" --http.addr 0.0.0.0 --http.api "eth,net,web3,personal,admin,miner" --allow-insecure-unlock --rpc.enabledeprecatedpersonal --syncmode "full" >> geth3.log 2>&1 &
```

## 连接小狐狸钱包
### 连接网络
RPC： http://xxx:8545
chainId: 12345
名称和代币符号任意。
### 导入节点
输入节点私钥导入节点账户

## 连接节点
```shell
geth attach /opt/etherData/node0/geth.ipc
```

1. 查看节点信息
```shell
admin.nodeInfo.enode
```
通过addPeer命令添加节点.
```shell
admin.addPeer("节点信息")
```

2. 根据节点私钥导入账号，提供节点私钥、加密节点私钥的对称密钥
```shell
personal.importRawKey("08a7533871d3a2e01d3a8849320cbfb703eb20c5dd2a9ccd2d9780eba5659c8e","yu201219jing")
```
3. 查看所有账户列表
```shell
eth.accounts
```
4. 查看所有账户余额
```shell
eth.getBalance(eth.accounts[0])
```
```shell
balanse=web3.fromWei(eth.getBalance(eth.accounts[0]),'ether')
```
5. 查询区块高度
```shell
eth.blockNumber
```
6. 查看矿工账户
```shell
eth.coinbase
```
7. 设置矿工账户
```shell
miner.setEtherbase(eth.accounts[0])
```
8. 启动挖矿（start（） 的参数表示挖矿使用的线程数）/关闭挖矿
```shell
miner.start()
```
```shell
miner.stop()
```
9. 交易操作
涉及链上交易时，需要先解锁账户。

解锁账户

personal.unlockAccount(address, passphrase, duration),密码和解锁时长都是可选的。如果密码为null，控制台将提示交互输密码。解密的密钥将保存在内存中直到解锁周期超时，默认的解锁周期为300秒。将解锁周期设置为0秒将解锁该密钥直到退出geth程序。
```shell
personal.unlockAccount(eth.accounts[0],'passward',0)
```

转账操作
```shell
eth.sendTransaction({from:eth.accounts[0],to:eth.accounts[1],value:web3.toWei(4,'ether')})
```

根据交易hash查询交易数据
```shell
eth.getTransaction("TxHash")
```

8. 出块权限
目前的验证节点通过发起提案增加出块节点，增加后的节点和当前的验证者轮流出块。
```shell
clique.propose("ADDR",true)
```

9. 关闭节点

```shell
ps -ef | grep geth
```
```shell
kill -15 PID
```

10. 清除链数据
```shell
geth removedb --datadir "/opt/etherData/node1"
```

11. 关闭链
>  ps aux | grep geth | grep -v grep | awk '{print $2}'| xargs kill -15