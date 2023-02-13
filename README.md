# jingixu cli web 项目模板
此项目仅为 jingxiu 脚手架提供项目模板

# 项目目录说明
```
jingxiu_initial_workspace
    - common        # 公用组件层
    - data          # 数据交互层
    - etc           # 配置信息
    - gateway       # 网络服务层
    - services      # 微服务层
```

# 分层说明
### 1. common 公用组件层
此层提供了一些通用的方法包，可共开发快速使用

active 包是一个简述用户活跃信息的组件

format 包是放置格式化信息的函数集合

jwt    包提供了生成token，刷新token，解析token的能力

network 包提供了网络请求能力

result  包封装了请求返回

utils 包提供了一些可用函数，有并行器，模拟器，一些类型转换及通用方法

### data 数据交互层
此层用于为开发着提供与数据库交互的能力

注：使用命令生成 【jingxiu model --src ...】

### gateway 网络服务层
此层是提供网络服务接口的管理层

config/config.go 读取配置文件中的可选信息

handle 是放置请求钩子函数的地方

router 初始化网络服务端，并挂载生成的路由文件

services 用于链接gRpc的客户端