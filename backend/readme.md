# 用golang写的一个nas系统
## 需求
* 后台文件管理，web
* 客户端所有平台都支持 web、win、mac、android、ios
* 账号管理

## 环境
* 数据库 mysql5.7
* golang 1.14
## 代码目录结构
主目录下的代码主要为后端代码，前端代码在fronted目录
common 一些常用函数的封装
routes 路由 遵循restful
fronted 所有前端代码
backend 所有后端代码

## golang后端功能实现
### 文件存储
* 文件不上传，以文件的md5为准
* 文件不按目录存储，目录只是文件的一个属性而已

## 后台管理系统


## 客户端
### web
## 账号
角色
## restful规范

状态：
* 200 OK 服务器返回用户请求的数据，该操作是幂等的
* 201 CREATED 新建或者修改数据成功
* 204 NOT CONTENT 删除数据成功
* 400 BAD REQUEST 用户发出的请求有问题，该操作是幂等的
* 401 Unauthoried 表示用户没有认证，无法进行操作
* 403 Forbidden 用户访问是被禁止的
* 422 Unprocesable Entity 当创建一个对象时，发生一个验证错误
* 500 INTERNAL SERVER ERROR 服务器内部错误，用户将无法判断发出的请求是否成功
* 503 Service Unavailable 服务不可用状态，多半是因为服务器问题，例如CPU占用率大，等等

url:
* GET /collections 返回资源列表
* GET /collections/:id 返回单独的资源
* POST /collections 返回新生成的资源对象
* PUT /collections/:id 返回完整的资源对象
* PATCH /collections/:id 返回被修改的属性
* DELETE /collections/:id 返回一个空文档
