# PrivatePOAEthereum

1. 安装golang gcc

> sudo apt install build-essential

2. Clone Ethereum Code

> git clone https://github.com/ethereum/go-ethereum.git

3. 编译geth工具

> cd go-ethereum && make geth

4. 写入geth的环境变量
```shell
vi ~/.bashrc

export ETHPATH=/opt/go-ethereum/build/bin
export PATH=$ETHPATH:$PATH
 
source ~/.bashrc
```

5. 生成账户地址 

私钥通过用户输入的密码加密存储,可以通过
> geth account new --datadir ./node1