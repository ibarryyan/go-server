# count_num

基于Go语言开发简单后端业务系统

### 1 技术点：

框架：

- Gin
- GORM
- protobuf

数据库：

- MySQL
- Redis

插件：

- cast
- Viper

### 2 包含功能

- 增删改查
- 分页

### 3 接口规范

#### 3.1 HTTP接口

POST("/add/:key")

GET("/findByKey/:key")

GET("/findById/:id")

POST("/saveInfo")

POST("/deleteInfo/:id")

GET("/getAll")

POST("/update")


#### 3.2 RPC接口


```go
//service接口
service NumInfoService {
  //rpc接口中的方法
  rpc GetNumInfoById(InfoRequest) returns (InfoResponse){}
  rpc AddNumByKey(InfoRequest) returns (InfoResponse){}
  rpc FindNumInfoByKey(InfoRequest) returns (InfoResponse){}
  rpc SaveNumInfo(InfoRequest) returns (InfoResponse){}
  rpc DeleteById(InfoRequest) returns (InfoResponse){}
  rpc FindAll(InfoRequest) returns (InfoResponse){}
}

//请求的数据格式 message 对应生成的代码中的struct，[修饰符] 类型 字段名 = 标识符
message InfoRequest{
  int64 id = 1;
  string name = 2 ;
  string info_key = 3 ;
  int64  info_num = 4;
}

message InfoResponse{
  int64 code = 1;
  string msg = 2;
  int64 count = 3;
  string data = 4;
}
```


### 4 相关博客：

- [《从0开始，用Go语言搭建一个简单的后端业务系统》](https://blog.csdn.net/Mr_YanMingXin/article/details/125294855)

- [《从1开始，扩展Go语言后端业务系统的RPC功能》](https://blog.csdn.net/Mr_YanMingXin/article/details/125317457)

- [《从2开始，在Go语言后端业务系统中引入缓存》](https://blog.csdn.net/Mr_YanMingXin/article/details/125365686)

- [《从3开始，在业务系统中增加分页功能》](https://blog.csdn.net/Mr_YanMingXin/article/details/125420590)