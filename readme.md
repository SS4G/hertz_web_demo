## 如何运行
1. 使用docker 创建mysql实例作为数据库
```bash
host_data_path="/Users/xxxx/Desktop/go_proj/website1/hertz_demo/hertz_mysql_data"
container_name="mysql_hertz"
password="123456"
docker run --name ${container_name} -e MYSQL_ROOT_PASSWORD=${password} -p 12001:3306 -v ${host_data_path}:/var/lib/mysql -d mysql:latest
```
2. ./build_run.sh 编译并运行服务
3. 在浏览器输入 127.0.0.1:8888/index 来访问网页效果如图
4. ![img.png](img.png)