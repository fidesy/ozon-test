# ozon-test
Решение тестового задания на позицию Golang стажера-разработчика.

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
        ├── errors/   
        └── persistence/
            ├── inmemory/         # in-memory решение
            ├── postgres/         # взаимодействие с PostgreSQL
            └── repository.go     # фабричный метод для создания базы данных
    └── usecase/
        ├── url.go                # реализация поведения сущности url
        └── usecase.go               
├── pkg/utils/
        ├── generate_short_url.go # функция генерация сокращенной ссылки
        └── load_config.go        # метод для загрузки .yml конфига
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

3. Выбрать необходимую базу данных [./configs/config.yaml](./configs/config.yaml#3) 
поле *database*, опции: postgres, in-memory

4. Запустить приложение.
```
docker compose up -d
```

### Использование 

### REST API
Создание сокращенной ссылки
```bash
curl -X POST -d '{"original_url": "https://ozon.ru"}' "http://localhost:8000/api/create"

# {"short_url":"http://localhost:8000/WkLSvXb94k"}
```

Получение оригинальной ссылки
```bash
curl http://localhost:8000/WkLSvXb94k

# {"original_url":"https://ozon.ru"}
```