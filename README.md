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
