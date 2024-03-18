# chat-mannual
Sprint:
https://www.notion.so/chat_manual-f79fd2f1542540f5b96515032bcb2d5b

Create a mysql server instance:
 docker run -itd --name <container_name> -p 3306:3306 -e MYSQL_ROOT_PASSWORD=<Yourpassword> MYSQL_DATABASE=sqldb -d

Connect to mySQL container:
jdbc:mysql://localhost:<port>/<dbname>?allowPublicKeyRetrieval=true&useSSL=false