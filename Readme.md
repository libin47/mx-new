# ubuntu 22.04
### apt使用清华镜像源
```bash
sudo sed -i "s@http://.*archive.ubuntu.com@https://mirrors.tuna.tsinghua.edu.cn@g" /etc/apt/sources.list
sudo sed -i "s@http://.*security.ubuntu.com@https://mirrors.tuna.tsinghua.edu.cn@g" /etc/apt/sources.list
sudo apt update
```

### git/curl/vim安装
```bash
sudo apt install git
sudo apt install curl
sudo apt install vim
```
添加ssh密钥
```bash
ssh-keygen -t rsa -C "your@email.com"
cat ~/.ssh/id_rsa.pub
```
clone it to github/gitlab'ssh config
### vscode安装
```bash
wget https://az764295.vo.msecnd.net/stable/e2816fe719a4026ffa1ee0189dc89bdfdbafb164/code_1.75.0-1675266613_amd64.deb
```
或者去官网上下最新版,然后安装
```bash
sudo dpkg -i code_****_amd64.deb
```
# database
### mongodb
```bash
wget -qO - https://www.mongodb.org/static/pgp/server-6.0.asc | sudo apt-key add -

echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/6.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-6.0.list

sudo apt update

curl -LO http://archive.ubuntu.com/ubuntu/pool/main/o/openssl/libssl1.1_1.1.1-1ubuntu2.1~18.04.20_amd64.deb
sudo dpkg -i ./libssl1.1_1.1.1-1ubuntu2.1~18.04.20_amd64.deb

sudo apt install -y mongodb-org
```
start
```bash
sudo systemctl start mongod  # start mongodb
sudo service mongod start # or this
```
look status
```bash
sudo systemctl status mongod  # check status of mongodb
```
### redis
```bash
sudo apt install -y redis
```
start
```bash
sudo systemctl start redis
sudo service redis start
```

# clash
教程 : https://juejin.cn/post/7127911250654134302    
注意有无窗口版本的区别


# nodejs
### nvm
```bash
wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh | bash
```
if can't connect, download and run
```bash
wget https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.3/install.sh
bash install.sh
```

```bash
source ~/.bashrc
# 切换源
nvm npm_mirror https://npm.taobao.org/mirrors/npm/
```
### nodejs
```bash
nvm install 19.5.0 / 16.15(server)
# pnpm
npm install pnpm -g
# pnpm源
pnpm config set registry https://registry.npm.taobao.org
```

# go
```bash
wget https://go.dev/dl/go1.19.5.linux-amd64.tar.gz
sudo  rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.19.5.linux-amd64.tar.gz
```
add to system PATH

add ```export PATH=$PATH:/usr/local/go/bin``` to /etc/profile    

添加代理
```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn
```
初始化&安装依赖
```
go mod init
go get
```

# start-old
```bash
git clone https://github.com/libin47/mx-server.git
git clone https://github.com/libin47/mx-kami.git
git clone https://github.com/libin47/mx-admin.git
```
在各个目录下
```bash
git checkout new
```

## mx-server
test:
node 16.15
```
pnpm i
npm run dev
```



# start-new


