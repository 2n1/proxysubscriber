# proxy subscriber

构造你自己的代理订阅服务。

> go + gin + sqlite + sb admin2

## 支持的协议

+ vmess + tls + websocket

+ trojan

+ Shadowsocks

## 支持的客户端

+ clash 系

+ v2ray 系

## CF优选IP

对于vmess协议，支持CF面向三大运营商的优选IP

## 演示

[ps.i2n1.cf](https://ps.i2n1.cf/login)

登录信息

```
Email: 2n1@i2n1.cf
密码：i2n1.cf
```

## 安装

1. 下载对应版本的预编译二进制包

2. 将 `config.json.example` 重命名为 `config.json`

3. 根据你的需要修改 `config.json` 的内容

|配置项|说明|
|----|----|
|`addr`|监听地址|
|`page_size`|分页大小|
|`mode`|运行模式，可选值：`debug`和`release`|
|`db_file`|SQLite数据库文件的路径|
|`sql_file`|创建表的SQL文件路径|
|`site_name`|站点名称|
|`base_url`|生成的订阅链接的前缀|
|`is_demo`|是否是演示站点|
|`session.secret_key`|Session的密钥|
|`session.name`|Session的名称|

4. 运行主程序
