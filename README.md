# chat-mannual
Sprint:
https://www.notion.so/chat_manual-f79fd2f1542540f5b96515032bcb2d5b

Mysql running in docker: https://medium.com/towards-data-engineering/dockerize-your-databases-a-step-by-step-guide-to-mysql-containerization-8dc2deabeebd

Create a mysql server instance:
 docker run --name <contianername> -p 3306:3306 -e MY_SQL_ROOT_PASSWORD=root -d MYSQL_DATABASE=sqldb mysql

Connect to mySQL container:
jdbc:mysql://localhost:<port>/<dbname>?allowPublicKeyRetrieval=true&useSSL=false

Run Msql server inside the container:
'docker exec -it my-sql-server /bin/bash'

To run our mysql instance inside the container, we would simply run the following command in terminal:

'mysql -uroot -p -A'

--u指定用户名为root
--p表示需要输入密码进行登录
--A表示使用全局选项，即启用全局模式并禁用不安全的语句。

