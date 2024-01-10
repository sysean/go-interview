# 使用官方 MySQL 镜像作为基础镜像
FROM mysql:latest

# 设置环境变量，用于配置 MySQL 服务器
ENV MYSQL_DATABASE students_database
ENV MYSQL_USER sean
ENV MYSQL_PASSWORD abc123
ENV MYSQL_ALLOW_EMPTY_PASSWORD yes

# 将初始化学生表的 SQL 脚本复制到容器中
COPY init.sql /docker-entrypoint-initdb.d/

# 将自定义的 my.cnf 配置文件复制到容器中，配置了binlog
COPY my.cnf /etc/mysql/conf.d/

# 对外暴露 MySQL 服务器的默认端口
EXPOSE 3306
