## goblog 

goblog 是基于beego框架开发的博客系统。主张简约、简单。系统采用mongodb数据库存储，七牛云存储提供cdn支持，markdown书写文章，拥有完整后台，可docker部署。各位朋友，不要吝啬你的star，赏个星星吧。

#### 快速开始

下载安装
```
go get github.com/deepzz0/goblog
```

安装mongodb数据库
```
brew install mongodb
```
若没有<code>brew</code>，可自行谷歌，安装。
配置mongodb地址，mongodb默认读取环境变量<code>MGO</code>
你可以到修改<code>~/.bash_profile</code>，添加如下
```
export MGO="127.0.0.1"
```
#### 配置文件

<code>conf/app.conf</code>
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
<code>runmode</code>选择运行模式，dev会输出beego日志，监听8080端口，主机名为127.0.0.1
配置域名地址<code>mydomain</code>，该域名相关各个链接地址。

<code>17monipdb.dat</code>
该文件用于后台统计ip地理位置用，当然你也可以直接采用其它的统计方式。

<code>backleft</code>
该文件用于配置后台菜单，请不要删除修改。

<code>qiniu.conf</code>
```
#------------- cdn -------------
accesskey = MB6AXl_Sj_mmFsL-Lt59Dml2Vm****
secretkey = BIrMy0fsZ0_SHNceNXk3e***
bucket = goblog
domain = 7xokm2.**.**.clouddn.com
zone = 0
```
该文件为七牛cdn，配置之后你在写文章时，可以直接上传文件倒cdn。

<code>backup</code>
该文件是前台展示，账号信息的模版，你可以直接修改配置。或者在程序运行成功后，后台修改。
``` json
{
    "UserName": "deepzz",
    "PassWord": "deepzz",
    "Salt": "__(f",
```
上面配置你后台的用户名，密码，随机盐。配置时使用的是明文，数据库存储是加密过的。  
注意，你需要到<code>models/model.go</code>修改默认用户，将deepzz替换成你的用户名。
``` go
	UMgr.loadUsers()
	Blogger = UMgr.Get("deepzz")
```

#### 多说评论框架
``` js
<!-- 多说评论框 start -->
    <div class="ds-thread" data-order="desc" data-limit="20" data-form-position="top" data-thread-key="{{.ID}}" data-title="{{.Title}}" data-url="{{$.Domain}}/{{.URL}}"></div>
    <!-- 多说评论框 end -->
    <!-- 多说公共JS代码 start (一个网页只需插入一次) -->
    <script type="text/javascript">
      var duoshuoQuery = {short_name:"deepzz"};
      (function() {
        var ds = document.createElement('script');
        ds.type = 'text/javascript';ds.async = true;
        ds.src = (document.location.protocol == 'https:' ? 'https:' : 'http:') + '//static.duoshuo.com/embed.js';
        ds.charset = 'UTF-8';
        (document.getElementsByTagName('head')[0] || document.getElementsByTagName('body')[0]).appendChild(ds);
      })();
    </script>
    <!-- 多说公共JS代码 end -->
```
你需要到多少官网获取的你网站的shot_name，将上面的deepzz替换掉。该代码段嵌到多个页面，你需要一一替换。

#### 统计相关
``` js
<script>
  (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
  (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
  m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
  })(window,document,'script','https://www.google-analytics.com/analytics.js','ga');

  ga('create', 'UA-77251712-1', 'auto');
  ga('send', 'pageview');

</script>
```
这里博主用了Google的数据统计，你可以将上面的代码替换成百度的统计代码等其它统计方式。在<code>views/homelayout.html</code>

#### 插件相关
UserAgent Parser，基于GO的用户代理解析器。可以到<code>domain:port/plugin/useragent.html</code>访问。

#### 有关其它
<code>static</code>目录下:

1. <code>feedTemplate.xml</code>是生成feed.xml的模版，你可以通过访问<code>domain:port/feed</code>查看，每小时自动更新。
2. <code>robots.txt</code>，网络爬虫排除协议。
3. <code>sitemap.xml</code>，网站地图，用于搜索引擎快速收录，博主爬虫尚未写好，现只能通过后台手动配置，你也可以通过自己的方式处理。访问两种方式<code>domain:port/sitemap</code>和<code>domain:port/sitemap.xml</code>。

所有都配置完成，在根目录下运行<code>bee run</code>

#### 展示
可以到我的博客[http://deepzz.com](http://deepzz.com)查看，国外服务器网速稍慢。  
前端页面
![show](http://7xokm2.com1.z0.glb.clouddn.com/img/home.png)
后台登陆
![login](http://7xokm2.com1.z0.glb.clouddn.com/img/login.png)
首页统计
![analysis](http://7xokm2.com1.z0.glb.clouddn.com/img/analysis.png)
博文修改
![modify](http://7xokm2.com1.z0.glb.clouddn.com/img/modify.png)



