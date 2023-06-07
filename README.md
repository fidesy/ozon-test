# ozon-test
Решение тестового задания на позицию Golang стажера-разработчика.

### Результат работы

---
Сервер принимает два запроса:

**GET /:hash** - получение оригинальной ссылки по сокращенной

Возможные ответы:
    
* 200 (OK) - хэш успешно найден, оригинальная ссылка возвращена
* 404 (NotFound) - хэш не найден
* 500 (InternalServerError) - ошибка сервера


**POST /api/create** - создание сокращенной ссылки

Возможные ответы:

* 201 (Created) - сокращенная ссылка успешно создана
* 400 (BadRequest) - клиент предоставил невалидное тело запроса / ссылку
* 409 (Conflict) - данная ссылка уже существует
* 500 (InternalServerError) - ошибка сервера

---
**_Алгоритм генерации сокращенных ссылок_**
1. Вычисление хэш-суммы MD5 исходной ссылки.
2. Преобразование хэша в строку с использованием кодирования base64.
3. Замена знаков (/+) на _ в полученной строке по условию задачи.
4. Строка обрезается до длины в 10 символов.

Что насчет коллизий?

В случае возникновения коллизий используется простой алгоритм - добавление вычисленного хэша к исходной ссылке в цикле пока не найдется уникальный хэш. Довольно наивный алгоритм, но самое главное, что выполняется условие идемпотентности.

---

### Структура проекта

```
├── .github/             
     └── workflows/
         └── main.yml             # конфигурационный файл для Github Actions
├── cmd/ozon-test/
         └── main.go              # entrypoint - запуск приложения
├── configs/config.yml            # конфигурационные параметры сервиса
├── internal/
    ├── config/config.go          # структуры для хранения конфига    
    ├── delivery/
        ├── grpc/                 # grpc сервис 
        └── handler/              # rest api 
    ├── domain/url.go             # описание структуры и поведения сущности
    ├── infrastructure/
        ├── dberrors/   
        └── persistence/
            ├── inmemory/         # in-memory решение
            ├── postgres/         # взаимодействие с PostgreSQL
            └── repository.go     # фабричный метод для создания баз данных
    └── service/
        ├── url.go                # реализация поведения сущности url
        └── service.go               
├── pkg/utils/
        └── generate_short_url.go # алгоритм генерации сокращенной ссылки
├── go.mod                        # зависимости проекта    
├── migrations                    # схемы таблиц базы данных
└── Makefile                     
```

### Установка

1. Клонировать репозиторий.
```
git clone https://github.com/fidesy/ozon-test.git
cd ozon-test
```

2. Создать файл .env, содержащий переменные для создания контейнера базы данных с помощью docker-compose
```
cp .env.example .env
``` 

3. Выбрать необходимую базу данных [./configs/config.yaml](./configs/config.yml#4) 
поле *database*, опции: postgres, in-memory. Также есть возможность переопределить данное значение указав соответствующих флаг при запуске сервиса в Dockerfile
```dockerfile
ENTRYPOINT ["main", "-db", "postgres"]
```

4. Запустить приложение.
```
docker compose up -d
```

### Использование 

**REST API**

Создание сокращенной ссылки
```bash
curl -i -X POST -d '{"original_url": "https://ozon.ru"}' "http://localhost:8000/api/create"

# Headers
# HTTP/1.1 201 Created
# Content-Type: application/json; charset=utf-8
# Date: Wed, 07 Jun 2023 12:26:01 GMT
# Content-Length: 48
# ----------------------------------
# Response body
# {"short_url":"http://localhost:8000/WkLSvXb94k"}
```

Получение оригинальной ссылки
```bash
curl -i http://localhost:8000/WkLSvXb94k

# Headers
# HTTP/1.1 200 OK
# Content-Type: application/json; charset=utf-8
# Date: Wed, 07 Jun 2023 12:26:39 GMT
# Content-Length: 34
# ----------------------------------
# Response body
# {"original_url":"https://ozon.ru"}
```

Повторное создание сокращенной ссылки для уже созданной оригинальной
```bash
curl -i -X POST -d '{"original_url": "https://ozon.ru"}' "http://localhost:8000/api/create"

# Headers
# HTTP/1.1 409 Conflict
# Content-Type: application/json; charset=utf-8
# Date: Wed, 07 Jun 2023 12:27:05 GMT
# Content-Length: 32
# ----------------------------------
# Response body
# {"message":"URL already exists"}
```

Попытка создать ссылку с невалидным телом запроса
```bash
curl -i -X POST -d '{"example": "https://ozon.ru"}' "http://localhost:8000/api/create"

# Headers
# HTTP/1.1 400 Bad Request
# Date: Wed, 07 Jun 2023 12:28:51 GMT
# Content-Length: 106
# Content-Type: text/plain; charset=utf-8
# ----------------------------------
# Response body
# {"message":"Key: 'URL.OriginalURL' Error:Field validation for 'OriginalURL' failed on the 'required' tag"}

curl -i -X POST -d '{"original_url": "http//invalid.com"}' "http://localhost:8000/api/create"

# Headers
# HTTP/1.1 400 Bad Request
# Content-Type: application/json; charset=utf-8
# Date: Wed, 07 Jun 2023 12:30:13 GMT
# Content-Length: 65
# ----------------------------------
# Response body
# {"message":"the URL provided is invalid and cannot be processed"}
```

Попытка получить ссылку по несуществующей сокращенной
```bash
curl -i http://localhost:8000/WkLSvXbOOO

# Headers
# HTTP/1.1 404 Not Found
# Content-Type: application/json; charset=utf-8
# Date: Wed, 07 Jun 2023 12:31:43 GMT
# Content-Length: 27
# ----------------------------------
# Response body
# {"message":"URL not found"}
```