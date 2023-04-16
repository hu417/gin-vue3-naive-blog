


# 简易博客社区后端

## 项目准备
相关工具:
    go      1.19.3
    vscode  go编辑器
    apifox  接口工具
    mysql   8.0.32


### 创建项目
```bash
mkdir -p go-vue-blogCommunity/blog_server
cd go-vue-blogCommunity/blog_server

go mod init blog_server
GOPROXY=https://goproxy.io,direct

```
### 模块安装
```bash
# gin
go get -u github.com/gin-gonic/gin

# gorm
go get -u github.com/jinzhu/gorm

# mysql驱动
go get -u github.com/go-sql-driver/mysql

# jwt
go get -u github.com/dgrijalva/jwt-go

# uuid
go get -u -v github.com/satori/go.uuid

```

### 创建数据库
```bash
]# mysql -uroot -p123456
mysql> create database `blog-community`;
```

### 创建静态文件目录
```bash
$ mkdir -p static/img
$ cp xx.jop static/img/
// xx.job 作为用户的初始头像
```

## 服务api接口

### 登录注册接口
#### 用户模型
```bash
$ mkdir -p model  // 存放项目的数据结构模型
$ touch model/user.go  // 用户结构体User
$ touch model/array.go // 自定义数组类型，并在数据存取时进行格式转换，即将数据存到数据库时，对数据进行处理，获得数据库支持的类型，而从数据库读取数据后，对其进行处理，获得Go类型的变量

```
#### 连接数据库
```bash
$ mkdir -p common/  // 该文件夹存放项目的一些通用功能
$ touch common/database.go  // 编写数据库初始化函数InitDB()与数据库数据获取函数GetDB()
```
#### 注册功能
```bash
$ mkdir -p cotroller  // 该文件夹存放主要的操作函数
$ touch controller/UserController.go  // 编写与用户有关的函数

```
#### jwt token
```bash
# 用户登录成功后需要为他发放一个token，前端接收到返回的token后会将其保存; 当请求需要token验证的接口时再发送给后端，此时，后端就需要对token进行解析，识别出用户的身份

$ touch common/jwt.go
```
##### token生成
// common/jwt.go

##### token解析
// common/jwt.go

#### 登录功能
```bash
# 用户登录函数的流程包括获取参数、数据验证、判断密码是否正确、发放token、返回结果，

// controller/UserController.go
```

#### 中间件
##### token验证
> 主要作用: 获取到前端请求中的token，调用ParseToken()对其进行解析，若token不合规范，该请求将会被抛弃，当token符合规范时才可以进行下一步操作
```bash
$ mkdir -p middleware
$ touch middleware/AuthMiddleware.go  // token验证
```
##### CORS跨域
> 前后端交互时,不同域名之间的请求需要处理跨域请求
```bash
$ touch middleware/CORSMiddleware.go
```

#### 用户信息
> 对前端发送的token进行解析并返回用户的部分信息
```bash

// controller/UserController.go  // 获取登陆后的用户信息

```
#### 编写路由
```bash
$ mkdir -p routes  // 路由文件夹
$ touch routes/routes.go

```
#### 入口文件
> 项目启动文件
```bash
$ touch main.go

```
#### 接口测试
```bash
$ go mod tidy   // 安装依赖
$ go run main.go  // 运行项目

```

### 图片上传接口
#### 上传图像功能
```bash
$ touch controller/FileController.go  // 该函数接收前端传来的图片文件，保存于后端的静态文件夹并返回图片url

```
#### 创建路由
```bash
// router/routes.go
```
#### 接口测试
```bash
$ go mod tidy   // 安装依赖
$ go run main.go  // 运行项目
```

### 分类查询接口
#### 分类模型
```bash
$ touch model/category.go 
```
#### 迁移数据表
```bash
// common/database.go

```
#### 响应封装
> 为项目封装一个统一的失败与成功的返回格式
```bash
$ mkdir -p response
$ touch response/response.go
```
#### 查询分类
```bash
$ touch controller/CategoryController.go // 实现对文章进行分类
```
#### 修改路由
```bash
// routes/routes.go
```
#### 接口测试
```bash
$ go mod tidy   // 安装依赖
$ go run main.go  // 运行项目
```

### 文章增删查改接口
#### 文章模型
```bash
$ touch model/time.go  // 将时间戳转化为实际时间
$ touch model/article.go  // 写入数据库存储的文章数据结构以及返回的文章数据结构

```
#### 迁移数据表
```bash

```
#### 文章数据类型
> 定义一个请求数据时的文章数据类型，保留文章增删查改操作的基本字段，以便后端接整个结构体并验证
```bash
$ mkdir -p vo
$ touch vo/article.go
```
#### 文章增删改查功能
```bash
// 文章增删改查的操作函数，其中List()函数返回关键字和分类查询的结果，count为满足关键字和分类的数据条数，便于前端进行分页
$ touch controller/ArticleController.go
```
#### 修改路由
```bash
// router/router.go
```
#### 接口测试
```bash
$ go mod tidy   // 安装依赖
$ go run main.go  // 运行项目
```

### 用户信息管理接口
#### 获取简要信息
> 在显示文章时需要同时显示作者头像，因此我们写一个函数返回文章作者的头像、文章作者的ID以及当前登录用户的ID，以便判断文章的作者是否是登录用户
```bash
// controller/UserController.go
```
#### 获取详细信息
> 在用户信息的详情页，需要展示用户的头像、用户名、文章列表、收藏夹、关注列表等信息，我们写一个函数返回上述信息
> 由于收藏夹、关注列表是自定义类型，我们需要将其转化为字符串数组才能使用IN查询
```bash
// controller/UserController.go
```
#### 修改信息功能
> 用户允许修改头像和用户名
```bash
// controller/UserController.go
```
#### 修改路由
```bash
// router/router.go
```
#### 接口测试
```bash
$ go mod tidy   // 安装依赖
$ go run main.go  // 运行项目

// 注意token信息与用户表对应的id
```
### 收藏关注接口
#### 收藏功能
> 查询登录用户是否有收藏当前的文章，并编写函数实现将文章ID添加到用户收藏夹和移除用户收藏夹
```bash
$ touch controller/CollectsController.go
```
#### 关注功能
> 关注功能的实现与收藏功能的实现类似，但要修改文章作者的粉丝数
```bash
$ touch controller/FollowingController.go
```
#### 修改路由
```bash
// router/router.go
```
#### 接口测试
```bash
$ go mod tidy   // 安装依赖
$ go run main.go  // 运行项目

// 注意token信息与用户表对应的id
```


