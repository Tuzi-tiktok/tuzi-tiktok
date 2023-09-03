# Tuzi-Tiktok
🐯 抖音项目

### 初步架构图

![](https://ivresse.top/api/img/2023-08-02-14-02-23945--models.png)
### 技术选型
- 开发框架
    - HTTP框架：Hertz
    - RPC框架：Kitex
    - ORM框架：GORM
- 数据存储
    - 数据库：MySQL
    - 对象存储：Minio
    - 缓存：Redis
- 服务治理
    - 注册中心：Nacos
    - 日志聚合：Grafana
- 部署
    - 反向代理：Nginx
    - 负载均衡：HAProxy
    - TLS证书管理：acme.sh
- 其他
    - 用户鉴权：JWT
    - 媒体处理：ffmpeg
    - 配置管理：Viper

```text
├── bin
│   ├── dumps   -- Debug 模式下对请求和响应进行Dump的所在目录
│   ├── logs    -- 服务日志
│   └── pid     -- 服务进程Pids
├── dockers     -- Docker 部署相关文件
├── idl         -- 接口定义文件所在目录
│   ├── hertz
│   └── kitex
├── internal    --  内部服务
│   ├── gateway   --- 应用网关服务
│   │   └── biz     --- 业务相关
│   │    
│   └── service   --- 内部微服务对应目录
│       ├── auth
│       ├── comment
│       ├── favorite
│       ├── feed
│       ├── filetransfer
│       ├── message
│       ├── publish
│       └── relation
└── shared     --- 服务基础设施层
    ├── config  --- 服务中心化配置
    │   └── test
    ├── dao     --- 数据访问对象层 整合GORM-Gen
    │   ├── generate
    │   ├── model
    │   ├── query
    │   └── test
    ├── kitex  --- Kitex 根据IDL生成的公共服务依赖
    │   └── kitex_gen
    ├── logger --- 整合Zap 日志模块
    ├── oss    --- 对象存储服务 仿SPI机制
    │   ├── internal
    │   ├── lfs
    │   ├── minio
    │   └── test
    ├── redis  --- Redis相关引入
    ├── secret --- Auth鉴权相关
    │   └── test
    └── utils  --- 公共依赖/工具相关
        ├── changes
        ├── consts
        ├── ffmpeg  --- FFmepeg视频抽帧
        └── mapstruct
```