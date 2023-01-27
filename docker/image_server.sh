#docker run --name my-custom-nginx -v /host/path/nginx.conf:/etc/nginx/nginx.conf:ro -d nginx

/Users/bytedance/Desktop/go_proj/website1/docker/image_server.sh

HOST_IMG_PATH=/Users/bytedance/Desktop/go_proj/website1/resources
HOST_CONF_PATH=/Users/bytedance/Desktop/go_proj/website1/conf/nginx/nginx.conf
docker run --rm -it --name my-custom-nginx -p 8086:80 -v ${HOST_CONF_PATH}:/etc/nginx/nginx.conf -v ${HOST_IMG_PATH}:/home/ftpuser/www/ nginx /bin/bash
# nginx -t # 测试nginx配置是否ok
# nginx # 启动nginx服务
# 参考 https://blog.csdn.net/weixin_46202067/article/details/112406028
