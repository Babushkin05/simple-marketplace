
# Тестовое задание в ВК -  Simple Marketplace

**Simple Marketplace** — это микросервисное приложение на Go, реализующее базовую платформу объявлений с регистрацией, авторизацией и CRUD-операциями над объявлениями. Архитектура построена с использованием gRPC, Gin, PostgreSQL и Docker.

---

## Архитектура

```
                         +-----------------+
                         |     Swagger     |
                         |   (OpenAPI UI)  |
                         +--------+--------+
                                  |
                                  v
                         +--------+--------+
                         |   API Gateway   | 
                         |   (Gin + REST)  |      
                         +--------+--------+         
                                  |                  
            +---------------------+------------------+
            |                                         |                   
            v                                         v                   
+---------------------+                    +----------------+    
|   Auth Service      |                    |  Goods Service |   
|   (gRPC, Token JWT) |  <---------------- |  (gRPC CRUD)   |  
+---------------------+   jwt-validation.  +----------------+   
```

---


### Запуск через Docker Compose

```bash
docker-compose up --build
```

---

## Документация API

Swagger доступен по адресу:  
 **http://localhost:8080/swagger/index.html**

В UI можно выполнить:
- регистрацию (`/register`)
- логин (`/login`)
- создание объявления (`/ads [POST]`)
- получение списка объявлений с фильтрами и сортировкой (`/ads [GET]`)

---

## Тестирование

Примеры cURL-запросов:

### Регистрация

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"login": "user1", "password": "pass123"}'
```

### Логин

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"login": "user1", "password": "pass123"}'
```

### Создание объявления

```bash
curl -X POST http://localhost:8080/ads \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Macbook Pro",
    "description": "2021 model, M1",
    "image_url": "http://example.com/mac.jpg",
    "price": 1999.99
  }'
```

### Получение объявлений с фильтрами

```bash
curl -X GET "http://localhost:8080/ads?page=1&page_size=5&sort_by=price&sort_order=desc&min_price=1000&max_price=3000"
```


---

## 📚 Используемые технологии

- Go (Gin, gRPC)
- PostgreSQL
- JWT
- Swagger (swaggo/swag)
- Docker, Docker Compose
- Protocol Buffers

---

## 🔒 Авторизация

Аутентификация и авторизация реализованы через JWT. Для защищённых запросов необходимо передавать токен в заголовке:

```
Authorization: <your_token>
```





