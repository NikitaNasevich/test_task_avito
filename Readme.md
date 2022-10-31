# Запуск через Docker 

1 - При необходимости переопределить в файле Dockerfile перменную (ENV) DATABASE =
user_name:user_password@(address:port)/data_base_name

где:
- user_name - имя пользователя указанный в образе mysql
- user_password - пароль пользователя указанный в образе mysql
- address:port - адрес и порт по которому доступен контейнер mysql

sql файла для инициализации базы данных [init.sql](https://github.com/NikitaNasevich/test_task_avito/blob/master/db/mysql/init.sql)

2 - При необходимости переопределить в файле Dockerfile пременные (ENV) PORT и EXPOSE

3 - Командой docker build PATH сбилдить образ микросервиса
где PATH - путь в корневую папку проекта

4 - Если контейнер mysql еще не запущен -> запустить 

5 - Запустить образ микросервиса командой docker run IMAGE_ID
где IMAGE_ID - id образа микросервиса (при необходимости можно найти через команду docker images)

6 - Готово! 

# Запуск через Docker Compose

1 - При необходимости переопределить в файле docker-compose.yml перменные окружения DATABASE =
user_name:user_password@(address:port)/data_base_name

и 

PORT = номер порта

где:
- user_name - имя пользователя указанный в образе mysql
- user_password - пароль пользователя указанный в образе mysql
- address:port - адрес и порт по которому доступен контейнер mysql

2 - Командой docker-compose build PATH сбилдить образ микросервиса и зависимой базой данных,
где PATH - путь в корневую папку проекта

3 - Командой docker-compose up запустить образ

4 - Готово!