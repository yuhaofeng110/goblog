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
配置mongodb地址，mongodb默认读取环境变量<code>MGO</code>
你可以到修改<code>~/.bash_profile</code>，添加如下
```
export MGO="127.0.0.1"
```
## 配置文件

监听端口在<code>goblog/conf/app.conf</code>修改
```
appname = goblog
runmode = prod
[dev]
httpport = 8080
mydomain = http://127.0.0.1
[prod]
httpport = 80
mydomain = http://deepzz.com
[test]
httpport = 8888
```
<code>runmode</code>选择运行模式，dev会输出beego日志，监听8080端口
配置域名地址<code>mydomain</code>，该域名相关各个链接地址。

配置cdn,该cdn存储的是你的静态文件，如.js,.css,图片。
```
#------------- cdn -------------
accesskey = MB6AXl_Sj_mmFsL-Lt59Dml2Vm****
secretkey = BIrMy0fsZ0_SHNceNXk3e***
bucket = goblog
domain = 7xokm2.**.**.clouddn.com
zone = 0
```

所有都配置完成，在根目录下运行<code>bee run</code>

效果图，展示地址：<http://deepzz.com>
![show](http://7xokm2.com1.z0.glb.clouddn.com/img/blog.png)


