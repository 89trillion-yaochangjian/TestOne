
### 1、整体框架
整体功能的实现思路

解析ini配置文件获取端口号。析命令行参数传入json文件路径，整理格式且只保留士兵的id，战斗力，名字，稀有度，Cvc,解锁等级，重新保存为json文件，启动时解析该文件。使用gin开发服务，等待客户请求，根据传入参数匹配json文件中相应数据返回页面

### 2、目录结构
```
.
├── README.md
├── __pycache__
│   └── locustfile.cpython-39.pyc
├── app
│   └── http
│       ├── http
│       └── httpServer.go
├── conf
│   ├── app.ini
│   ├── new_Soldier.json
│   └── soldier.json
├── internal
│   ├── ctrl
│   │   └── soldierCrtrl
│   │       └── soldierCrtrl.go
│   ├── handler
│   │   └── soldierHandler
│   │       └── soldierHandler.go
│   ├── router
│   │   └── soldierRouter
│   │       └── soldierRouter.go
│   ├── service
│   │   └── soldierService
│   │       ├── soldierService.go
│   │       └── soldierService_test.go
│   └── utils
│       ├── IniUtil
│       │   └── IniUtil.go
│       ├── JsonUtil
│       │   └── JsonUtil.go
│       └── PathUtil
│           └── PathUtil.go
└── locustFile.py


```

#### 3. 代码逻辑分层




|层|文件夹|主要职责|调用关系|其他说明|
| ------------ | ------------ | ------------ | ------------ | ------------ |
|应用层 |app/http/httpServer.go  |服务器启动 |调用路由层工具层   |不可同层调用
|路由层 |internal/router/soldierRouter  |路由转发 | 调用工具层 控制层 被应用层   |不可同层调用
|控制层 |internal/ctrl/soldierCrtrl  |请求参数处理，响应 | 调用handler层 被路由层调用    |不可同层调用
|handler层 |internal/handler/soldierHandler  |处理具体业务 | 调用路由层service层，被控制层调用    |不可同层调用
|service层   |internal/service/soldierService  |处理业务逻辑 | 调用工具层，被handler层调用    |可同层调用
|工具层  |internal/utils  |文件处理 | 被路由层 ，service层调用  |不可同层调用
| 配置文件 |conf  |文件存储 | 不存在调用关系    |不可同层调用

#### 4.存储设计
读取的数据使用map存储，来源json配置文件，ID 为key string类型，value为士兵的结构体

####5. 接口设计
|   请求方法| 接口地址  |   请求参数|   请求响应(例如)|
| ------------ | ------------ | ------------ | ------------ |
|  Http GET |   http://127.0.0.1:8000/RarityFindByID| id  |  {"Rarity":"4"} |
|   Http GET|   http://127.0.0.1:8000/AtkFindByID|  id |  {"Atk":"900"} |
| Http GET  |  http://127.0.0.1:8000/SoldierFindALLCycUnlock | rarity,cvc,unlockArena  | {"id":"11002","Name":"Giant","UnlockArena":"2","Rarity":"3","Atk":"3300","Cvc":"1000"}|
| Http GET  |http://127.0.0.1:8000/SoldierFindByCyc   |  cvc | {"id":"15902","Name":"Spear Thrower","UnlockArena":"4","Rarity":"3","Atk":"73","Cvc":"1000"}  |
| Http GET  | http://127.0.0.1:8000/SoldierEachStage  |unlockArena   |  {"id":"11002","Name":"Giant","UnlockArena":"2","Rarity":"3","Atk":"3300","Cvc":"1000"} |

#### 6. 第三方库
1. "gopkg.in/ini.v1"
   https://github.com/go-ini/ini
   用于解析ini文件

2. "spf13/pflag"
   https://github.com/spf13/pflag
   获取命令行参数传入的文件路径
   
#### 7. 如何编译执行
go build

编译为可执行文件

./http --path /Users/yaochangjian/go/src/AnalyseFile/conf/soldier.json





