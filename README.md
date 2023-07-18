# PrivatePOAEthereum

1. 安装golang gcc

```shell
sudo apt install build-essential
```

2. Clone Ethereum Code

```shell
git clone https://github.com/ethereum/go-ethereum.git
```

3. 编译geth工具

```shell
 cd go-ethereum && make geth
```

4. 写入geth的环境变量
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

5. 生成账户地址 

私钥通过用户输入的密码加密存储,可以通过golang解析出账户地址和私钥： https://github.com/yuhuajing/PrivatePOAEthereum/blob/main/parsePrivateKey.go

```shell
geth account new --datadir ./node1
```
> 0x0B587FFD0BBa122fb5ddc19AD6eEcEB1D2dBbff7

```shell
geth account new --datadir ./node2
```
> 0x20b3ee0a86c4C26086Fc414D5E5b771F6abd7A3f
```shell
geth account new --datadir ./node3
```
> 0xe52d030C29aFb2fC470EB0FdF5c31287343004D2

6. 构造创世区块
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
      "clique":{
         "period":5,
         "epoch":300
      }
   },
   "difficulty":"1",
   "gasLimit":"800000000",
   "extradata":"0x00000000000000000000000000000000000000000000000000000000000000005534F5024146D16a5C1ce60A9f5a2e9794e3F9810000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
   "alloc":{
      "0x0B587FFD0BBa122fb5ddc19AD6eEcEB1D2dBbff7":{
         "balance":"1000000000000000000000"
      },
      "0x20b3ee0a86c4C26086Fc414D5E5b771F6abd7A3f":{
         "balance":"1000000000000000000000"
      },
      "0xe52d030C29aFb2fC470EB0FdF5c31287343004D2":{
         "balance":"1000000000000000000000"
      }
   }
}
```

7. 