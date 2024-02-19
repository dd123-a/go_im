概述
wenqianIm 是一个基于 Go 语言开发的即时通讯系统后端模块。该模块使用了一系列优秀的第三方库和工具，提供了强大的功能支持。

依赖
gin-gonic/gin：用于构建 Web 服务的 HTTP 框架。
go-playground/validator/v10：用于验证和解析请求数据。
go-redis/redis：用于与 Redis 数据库进行交互。
gorilla/websocket：提供了 WebSocket 协议的实现，用于支持实时通讯功能。
jinzhu/gorm：用于操作数据库的 ORM 框架。
onsi/ginkgo：Golang 编写的 BDD 测试框架（间接依赖）。
onsi/gomega：Golang 编写的匹配器库（间接依赖）。
sirupsen/logrus：用于日志记录的库。
mongo-driver：用于操作 MongoDB 数据库的官方驱动程序。
golang.org/x/crypto：提供了加密算法的实现。
ini.v1：用于读取 INI 格式配置文件的库。
yaml.v2：用于处理 YAML 格式数据的库。
