#host_data_path="/Users/xxxx/Desktop/go_proj/website1/hertz_demo/hertz_mysql_data"
#container_name="mysql_hertz"
#password="123456"
#docker run --name ${container_name} -e MYSQL_ROOT_PASSWORD=${password} -p 12001:3306 -v ${host_data_path}:/var/lib/mysql -d mysql:latest
#docker run --name ${container_name} -e MYSQL_ROOT_PASSWORD=${password} -p 12001:3306 -d mysql:latest 

# 访问命令
mysql -h127.0.0.1 -P12001 -uroot -p123456
