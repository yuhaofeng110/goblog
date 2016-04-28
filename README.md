## goblog 

goblog 是基于beego框架开发的博客系统。主张简约、简单。系统采用mongodb数据库存储，七牛云存储提供cdn支持，markdown书写文章，拥有完整后台，可docker部署。

## 快速开始

下载安装
```
go get github.com/deepzz0/goblog
```

安装mongodb
```
brew install mongodb
```
若没有<code>brew</code>，可自行谷歌，安装。

## 配置文件

监听端口在<code>goblog/conf/app.conf</code>修改
mongodb默认读取环境变量<code>MGO</code>
你可以到修改<code>~/.bash_profile</code>，添加如下
```
export MGO="127.0.0.1"
```
