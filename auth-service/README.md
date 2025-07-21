# 🛡️ Auth Service — gRPC сервис аутентификации для маркетплейса

Этот сервис является частью микросервисного приложения "Simple Marketplace" и отвечает за:
- регистрацию пользователей,
- аутентификацию по логину/паролю,
- проверку JWT-токенов.

## 📦 Стек
- Язык: Go
- gRPC
- JWT
- PostgreSQL

## 🚀 Интерфейс (gRPC)

### AuthService

```
service AuthService {
  rpc Register(RegisterRequest) returns (RegisterResponse);
  rpc Login(LoginRequest) returns (LoginResponse);
  rpc ValidateToken(ValidateTokenRequest) returns (ValidateTokenResponse);
}
```

---

### `Register`

Регистрирует нового пользователя.

**Request**
```json
{
  "login": "user123",
  "password": "securepass"
}
```

**Response**
```json
{
  "user_id": "uuid",
  "login": "user123"
}
```

---

### `Login`

Авторизует пользователя по логину и паролю. Возвращает JWT-токен.

**Request**
```json
{
  "login": "user123",
  "password": "securepass"
}
```

**Response**
```json
{
  "token": "<JWT>"
}
```

---

### ✅ `ValidateToken`

Проверяет валидность JWT токена. Используется другими сервисами для авторизации.

**Request**
```json
{
  "token": "<JWT>"
}
```

**Response**
```json
{
  "user_id": "uuid",
  "login": "user123",
  "valid": true
}
```

---

## Конфигурация

Конфигурация задаётся в `config/local.yaml`:

```yaml
server:
  grpc_port: 50051

database:
  host: localhost
  port: 5432
  user: auth_user
  password: auth_pass
  dbname: auth_db
  sslmode: disable

jwt:
  secret: "supersecretjwtkey1234567890"
  ttl: 60m
```

---

## Схема БД

Создаётся автоматически из `schema.sql` при старте сервиса:

```sql
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);
```

---

## Запуск (Docker)

`Dockerfile`  — в папке `deployments`.

---

## Структура сервиса

```
auth-service/
.
├── Makefile
├── README.md
├── api
│   ├── gen
│   │   ├── user.pb.go
│   │   └── user_grpc.pb.go
│   └── user.proto
├── cmd
│   └── main.go
├── config
│   └── local.yaml
├── deployments
│   └── Dockerfile
├── go.mod
├── go.sum
└── internal
    ├── auth
    │   ├── hash.go
    │   └── jwt.go
    ├── config
    │   └── config.go
    ├── db
    │   ├── db.go
    │   ├── init_db.go
    │   ├── posgres.go
    │   └── schema.sql
    ├── grpc
    │   ├── handler.go
    │   └── server.go
    └── service
        └── auth.go
```

---


