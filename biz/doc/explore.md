### 安装
- 安装hertz的时候 不要更改GOPATH 更改了可能会导致某些包用不了
- `hz new -module hz_demo` 创建一个新项目 新项目 新项目的modulename 自己指定为 hz_demo
- 编译 `go build -o hertz_demo && ./hertz_demo`
### 样例代码测试
[x] 路由注册 
[x] 动态模板路由
[] 静态文件上传
[x] html渲染
[x] 数据库ORM
[] cookie


### 安装
- grom `go get -u gorm.io/gorm`
- grom_driver `go get -u gorm.io/driver/mysql`
