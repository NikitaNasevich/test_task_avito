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

sql файла для инициализации базы данных [init.sql](https://github.com/NikitaNasevich/test_task_avito/blob/master/db/mysql/init.sql)

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

3 - Командой docker-compose up запустить образ, если mysql ранее не знапускался, дать время на инициальзацию БД.

4 - Готово!

# Методы:


# Метод начисления средств на баланс. Принимает id пользователя и сколько средств зачислить.

Метод проверяет данные на валидность, в случае если данные валидны производит операцию, иначе возвращает соответсвующую ошибку с кодом:
id пользователя должно быть целым числом > 0, данный id должен быть в БД.  
Количество средств также должно быть числом >0.

Примеры запросов/ответов:

Запрос:

POST /v1/addBalance HTTP/1.1
Host: 0.0.0.0:3000
Content-Type: application/json
Content-Length: 37

{
  "UserId": 1,
  "Balance": -1
}'


Ответ:

HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
Date: Mon, 31 Oct 2022 14:04:38 GMT
Content-Length: 46

{
  "Message": "Balance must be above 0",
  "Code": 3
}

----------------------------------------------------------------------------------------------------------

Запрос:

POST /v1/addBalance HTTP/1.1
Host: 0.0.0.0:3000
Content-Type: application/json
Content-Length: 40

{
  "UserId": 1,
  "Balance": "asd"
}

Ответ:

HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
Date: Mon, 31 Oct 2022 14:06:22 GMT
Content-Length: 98

{
  "Message": "json: cannot unmarshal string into Go struct field .Balance of type float64",
  "Code": 1
}

---------------------------------------------------------------------------------------------------------

Запрос:

POST /v1/addBalance HTTP/1.1
Host: 0.0.0.0:3000
Content-Type: application/json
Content-Length: 37

{
  "UserId": 1,
  "Balance": 10
}

Ответ:

HTTP/1.1 200 OK
Date: Mon, 31 Oct 2022 14:07:38 GMT
Content-Length: 0

<Response body is empty>



# Метод резервирования средств с основного баланса на отдельном счете. Принимает id пользователя, ИД услуги, ИД заказа, стоимость.

Метод проверяет данные на валидность, в случае если данные валидны производит операцию, иначе возвращает соответсвующую ошибку с кодом:
id пользователя должно быть целым числом > 0, данный id должен быть в БД. 
id услуги должно быть целым числом > 0, данный id должен быть в БД. 
id заказа должно быть целым числом > 0, данный id должен быть в БД. 
Количество средств также должно быть числом >0.

Запрос (id заказа 0 (недопустимое значение)):

POST /v1/reserveFunds HTTP/1.1
Host: 0.0.0.0:3000
Content-Type: application/json
Content-Length: 87

{
  "UserId": 1,
  "ServiceId": 1,
  "OrderServiceId": 0,
  "ReserveBalance": 10
}

Ответ:

HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
Date: Mon, 31 Oct 2022 14:10:10 GMT
Content-Length: 117

{
  "Message": "Key: 'OrderServiceId' Error:Field validation for 'OrderServiceId' failed on the 'required' tag",
  "Code": 1
}

---------------------------------------------------------------------------------------------------------

Запрос (не существующий id услуги):

POST /v1/reserveFunds HTTP/1.1
Host: 0.0.0.0:3000
Content-Type: application/json
Content-Length: 88

{
  "UserId": 1,
  "ServiceId": 25,
  "OrderServiceId": 1,
  "ReserveBalance": 10
}

Ответ:

HTTP/1.1 500 Internal Server Error
Content-Type: application/json; charset=utf-8
Date: Mon, 31 Oct 2022 14:11:56 GMT
Content-Length: 83

{
  "Message": "Error 3819: Check constraint 'customer_chk_1' is violated.",
  "Code": 104
}

---------------------------------------------------------------------------------------------------------

Запрос (все данные валидны):

POST /v1/reserveFunds HTTP/1.1
Host: 0.0.0.0:3000
Content-Type: application/json
Content-Length: 87

{
  "UserId": 1,
  "ServiceId": 1,
  "OrderServiceId": 1,
  "ReserveBalance": 10
}

Ответ:

HTTP/1.1 200 OK
Date: Mon, 31 Oct 2022 14:16:01 GMT
Content-Length: 0

<Response body is empty>

# Метод разрезервирования средств с основного баланса на отдельном счете. Принимает id пользователя, ИД услуги, ИД заказа, стоимость.

Метод проверяет данные на валидность, в случае если данные валидны производит операцию, иначе возвращает соответсвующую ошибку с кодом:
id пользователя должно быть целым числом > 0, данный id должен быть в БД. 
id услуги должно быть целым числом > 0, данный id должен быть в БД. 
id заказа должно быть целым числом > 0, данный id должен быть в БД. 
Количество средств также должно быть числом >0.

Запрос (id заказа = 0):

POST /v1/cancelReserveFunds HTTP/1.1
Host: 0.0.0.0:3000
Content-Type: application/json
Content-Length: 87

{
  "UserId": 1,
  "ServiceId": 1,
  "OrderServiceId": 0,
  "ReserveBalance": 10
}

Ответ:

HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
Date: Mon, 31 Oct 2022 15:12:15 GMT
Content-Length: 117

{
  "Message": "Key: 'OrderServiceId' Error:Field validation for 'OrderServiceId' failed on the 'required' tag",
  "Code": 1
}

---------------------------------------------------------------------------------------------------------

Запрос (все данные валидны):

POST /v1/cancelReserveFunds HTTP/1.1
Host: 0.0.0.0:3000
Content-Type: application/json
Content-Length: 87

{
  "UserId": 1,
  "ServiceId": 1,
  "OrderServiceId": 1,
  "ReserveBalance": 10
}

Ответ:

HTTP/1.1 200 OK
Date: Mon, 31 Oct 2022 15:15:16 GMT
Content-Length: 0

<Response body is empty>


# Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии.Принимает id пользователя, ИД услуги, ИД заказа, сумму.

Метод проверяет данные на валидность, в случае если данные валидны производит операцию, иначе возвращает соответсвующую ошибку с кодом:
id пользователя должно быть целым числом > 0, данный id должен быть в БД. 
id услуги должно быть целым числом > 0, данный id должен быть в БД. 
id заказа должно быть целым числом > 0, данный id должен быть в БД. 
Количество средств также должно быть числом >0.

Запрос (id пользователя с балансом = null):

POST /v1/acceptProfit HTTP/1.1
Host: 0.0.0.0:3000
Content-Type: application/json
Content-Length: 89

{
  "UserId": 2,
  "ServiceId": 1,
  "OrderServiceId": 1,
  "ReserveBalance": 6.05
}

Ответ:

HTTP/1.1 500 Internal Server Error
Content-Type: application/json; charset=utf-8
Date: Mon, 31 Oct 2022 14:19:03 GMT
Content-Length: 69

{
  "Message": "Error 1048: Column 'order_id' cannot be null",
  "Code": 106
}

---------------------------------------------------------------------------------------------------------

Запрос (id услиги = 0):

POST /v1/acceptProfit HTTP/1.1
Host: 0.0.0.0:3000
Content-Type: application/json
Content-Length: 89

{
  "UserId": 1,
  "ServiceId": 0,
  "OrderServiceId": 1,
  "ReserveBalance": 6.05
}

Ответ:

HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
Date: Mon, 31 Oct 2022 14:21:12 GMT
Content-Length: 107

{
  "Message": "Key: 'ServiceId' Error:Field validation for 'ServiceId' failed on the 'required' tag",
  "Code": 1
}

---------------------------------------------------------------------------------------------------------

Запрос (все данные валидны):

POST /v1/acceptProfit HTTP/1.1
Host: 0.0.0.0:3000
Content-Type: application/json
Content-Length: 87

{
  "UserId": 1,
  "ServiceId": 1,
  "OrderServiceId": 1,
  "ReserveBalance": 10
}

Ответ:

HTTP/1.1 200 OK
Date: Mon, 31 Oct 2022 14:25:12 GMT
Content-Length: 0

<Response body is empty>

# Метод получения баланса пользователя. Принимает id пользователя.

Метод проверяет данные на валидность, в случае если данные валидны возвращает id пользователя и его баланс, иначе возвращает соответсвующую ошибку с кодом:
id пользователя должно быть целым числом > 0, данный id должен быть в БД. 


Запрос:

POST /v1/getBalance HTTP/1.1
Host: localhost:3000
Content-Type: application/json

{
  "UserId": 0
}

Ответ:

HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
Date: Mon, 31 Oct 2022 10:29:09 GMT
Content-Length: 101

{
  "Message": "Key: 'UserId' Error:Field validation for 'UserId' failed on the 'required' tag",
  "Code": 1
}

----------------------------------------------------------------------------------------------------------

Запрос:

POST /v1/getBalance HTTP/1.1
Host: localhost:3000
Content-Type: application/json

{
  "UserId": "1"
}

Ответ:

HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
Date: Mon, 31 Oct 2022 10:16:07 GMT
Content-Length: 95

{
  "Message": "json: cannot unmarshal string into Go struct field .UserId of type int64",
  "Code": 1
}

----------------------------------------------------------------------------------------------------------

Запрос несуществующего пользователя в БД:

POST /v1/getBalance HTTP/1.1
Host: localhost:3000
Content-Type: application/json

{
  "UserId": 25
}

Ответ:

HTTP/1.1 500 Internal Server Error
Content-Type: application/json; charset=utf-8
Date: Mon, 31 Oct 2022 10:31:29 GMT
Content-Length: 41

{
  "Message": "UserId not found",
  "Code": 101
}

----------------------------------------------------------------------------------------------------------

Запрос:

POST /v1/getBalance HTTP/1.1
Host: localhost:3000
Content-Type: application/json

{
  "UserId": 1
}

Ответ:

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 31 Oct 2022 10:36:10 GMT
Content-Length: 27

{
  "UserId": 1,
  "Balance": 3.95
}