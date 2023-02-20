# Приложение для ведения списков дел/покупок

<!-- ToC start -->
# Содержание

1. [Описание задачи](#Описание-задачи)
1. [Реализация](#Реализация)
1. [Endpoints](#Endpoints)
1. [Запуск](#Запуск)
1. [Примеры](#Примеры)
<!-- ToC end -->

# Описание задачи

Разработать API для авторизации, аутентификации, создания, редактирования и удаления списков дел/покупок по id пользователя, а так же редактирования и удаления самих списков.

# Реализация

- Следование дизайну REST API.
- Подход "Чистой Архитектуры" и техника внедрения зависимости.
- Работа с фреймворком [gin-gonic/gin](https://github.com/gin-gonic/gin).
- Работа с СУБД Postgres с использованием библиотеки [sqlx](https://github.com/jmoiron/sqlx) и написанием SQL запросов.
- Конфигурация приложения - библиотека [viper](https://github.com/spf13/viper).
- Запуск бд из Docker.
**Структура проекта:**
```
.
├── pkg
│   ├── handler     // обработчики запросов
│   ├── service     // бизнес-логика
│   └── repository  // взаимодействие с БД
├── cmd             // точка входа в приложение
├── schema          // SQL файлы с миграциями
├── configs         // файлы конфигурации
```

# Endpoints
- POST   /auth/sign-up       - авторизация пользователя по имени, логину и паролю.
    - Тело запроса:
        - name               - имя пользователя.
        - login              - логин пользователя.
        - password           - пароль.
- POST   /auth/sign-in       - аутентификация пользователя по логину и паролю. Пользователь получает bearer токен сроком действия 24 часа.
    - Тело запроса:
        - login              - логин пользователя.
        - password           - пароль.

- POST   /api/lists          - создание списка дел.
    - Тело запроса:
        - title              - название.
        - description        - описание.
- GET    /api/lists          - получение всех списков дел.
- GET    /api/lists/id       - получение списка дел по id.
- PUT    /api/lists/id       - изменение списка дел по id.
    - Тело запроса:
        - title              - название.
        - description        - описание.
- DELETE /api/lists/id       - удаление списка дел по id.

- POST   /api/lists/id/items - создание дел внутри списка.
    - Тело запроса:
        - title              - название.
        - description        - описание.
- GET    /api/list/id/items  - получение дел списка.

- GET    /api/items/:id      - получение конкретного дела по id.
- PUT    /api/items/:id      - редактирование дела по id.
    - Тело запроса:
        - title              - название.
        - description        - описание.
        -  done              - статус выполнения.
- DELETE /api/items/:id      - удаление дела по id.

# Запуск

```
go run cmd/main.go
```

Если приложение запускается впервые, необходимо применить миграции к базе данных:

```
docker pull postgres
docker run --name=balance-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
go mod tidy
```


# Примеры

Запросы сгенерированы из Postman.

### 1. POST  /auth/sign-up
**Запрос:**
```
{
    "name" : "fedor",
    "username" : "fshmidt",
    "password" : "qwerty"
}
```
**Тело ответа:**
```
{"id":1}
```

### 2. POST  /auth/sign-in
**Запрос:**
```
{
    "username" : "fshmidt",
    "password" : "qwerty"
}
```
**Тело ответа:**
```
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzM5NzcxMzcsImlhdCI6MTY3MzkzMzkzNywidXNlcl9pZCI6M30.FueN9GuYJPgbcIJLPxxk6X0gkzb5QFh7faxF9Ch2Dak"}
```

### 3. POST   /api/lists
**Запрос:**
```
{
    "title" : "Купить",
    "description" : "до пятницы"
}
```
**Тело ответа:**
```
{"id": 1}
```
### 3.  GET  /api/lists
**Тело ответа:**
```
{
    "data": [
        {
            "id": 2,
            "title": "Купить",
            "description": "до пятницы"
        },
        {
            "id": 1,
            "title": "Продать",
            "description": "до субботы"
        }
    ]
}
```
### 4. GET  /api/lists/id
**Тело ответа:**
```
{
    "id": 2,
    "title": "Купить",
    "description": "до пятницы"
}

```

### 5. PUT   /api/lists/id
**Запрос:**
```
{
    "title" : "Продать",
    "description" : "до субботы"
}
```
**Тело ответа:**
```
{"status": "ok"}
```

### 6. DELETE /api/lists/id
**Тело ответа:**
```
{"status": "ok"}
```

### 7. POST   /api/lists/id/items

**Тело запроса:**
```
{
    "title" : "квартира",
    "description" : "3-комнатная"
}
```
**Тело ответа:**
```
{
    "id": 1
}
```

### 8. GET    /api/list/id/items
**Тело ответа:**
```
[
    {
        "id": 1,
        "title": "квартира",
        "description": "3-комнатная",
        "done": false
    },
    {
        "id": 2,
        "title": "магнитола",
        "description": "импортная",
        "done": false
    }
]
```

### 9. GET    /api/items/id
**Тело ответа:**
```
{
    "id": 1,
    "title": "квартира",
    "description": "3-комнатная",
    "done": false
}
```
### 10. PUT    /api/items/id
**Запрос:**
```
{
    "title": "квартира",
    "done": true
}
```
**Тело ответа:**
```
{
    "status": "ok"
}
```
### 10. DELETE  /api/items/id
**Тело ответа:**
```
{
    "status": "ok"
}
```
