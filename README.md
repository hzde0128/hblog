# hblog个人博客系统

Go语言开发的个人博客系统，基于Beego框架开发。

## 开发计划

- [ ] 后台模块
  - [x] 登录
  - [ ] Cookie,Session验证
  - [ ] 分类管理
  - [ ] 标签管理
  - [ ] 文章管理
  - [ ] Redis缓存
  - [ ] 系统设置
  - [x] 数据模型
  - [x] 基本路由设定
- 前台模块
  - [ ] 界面展示

## 安装

```bash
go get github.com/hzde0128/hblog
cd $GOPATH/github.com/hzde0128/hblog
go run main.go
```

## 登录

[http://127.0.0.1:8080/admin/login](http://127.0.0.1:8080/admin/login)

用户名: `admin`  
密码: `admin888`

