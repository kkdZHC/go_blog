## 页面
![img](https://github.com/kkdZHC/go_blog/blob/main/public/resource/images/example.png)
## 配置文件
```
[viewer]
    Title = "goblog"
    Description = "Go语言博客"
    Logo = "/resource/images/logo.png"
    Navigation = ["首页","/", "GO语言","/golang", "归档","/pigeonhole", "关于","/about"]
    Bilibili = "https://space.bilibili.com/399943874"
    Zhihu = "https://www.zhihu.com"
    Avatar = "https://b0.bdstatic.com/ugc/jsiDRNxzSfNyP64A4FO1LA959cad3cb2faa0c763017a102ce54755.jpg@h_1280"
    UserName = "你的名字"
    UserDesc = "好好学习天天向上"
[system]
    CdnURL = ""
    QiniuAccessKey = ""
    QiniuSecretKey = ""
    Valine = true
    ValineAppid = ""
    ValineAppkey = ""
    ValineServerURL = ""
```
## 目录结构
```
├─api
├─common
├─config
├─dao
├─models
├─public
│  ├─html
│  └─resource
├─router
├─service
├─template
│  └─layout
├─utils
└─views
```
