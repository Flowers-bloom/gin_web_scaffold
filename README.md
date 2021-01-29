# gin_web_scaffold
gin 初学者，本着边学习边实践的思想，致力于搭建一个简洁易用的基于 gin 的 web 脚手架。  

## 依赖
```
go mod init
go get github.com/gin-gonic/gin

go get github.com/spf13/viper

go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

go get github.com/go-redis/redis/v8
```

## 目录结构
- api 接口（控制）层  
- common 定义一些公共实现  
- config 配置层  
- constant 定义一些全局常量  
- dao 数据访问层  
- dto 数据传输层  
- model 数据实体层  
- router 定义路由信息  
- service 业务（服务）层  
- util 定义一些工具  

## 功能
1. gin 实现 web 路由、中间件等功能  
2. gorm 实现数据库访问  
3. viper 实现加载配置文件  
4. 通过 fmt 自定义实现日志框架，可同时写入控制台和文件，且支持文件滚动收集  
5. go-redis 封装对 redis 的操作  
6. 内部异常统一采用 panic 抛出，在上层函数中采用 defer、recover 捕获处理，避免程序终止  

## 日志
|    完成时间   |         功能点        |
| :---------: | :-------------------: |
| 2021.1.22 | gin 实现基础 pong 路由测试 |
| 2021.1.25 | 搭建自定义 web 目录结构，并实现 gin 路由 |
| 2021.1.26 | 依赖 gorm 实现数据库访问；依赖 viper 实现加载配置文件 |
| 2021.1.27 | 通过 fmt 自定义实现日志框架，可同时写入控制台和文件，且支持文件滚动收集 |
| 2021.1.28 | go-redis 封装字符串和一些通用操作，并统一异常处理，新增学生、学科、成绩三表，尝试原生 sql 查询 |
| 2021.1.29 | 结构体参数验证测试 |