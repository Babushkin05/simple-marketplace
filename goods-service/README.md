# Goods Service

Сервис управления объявлениями для платформы `simple-marketplace`.

## Функциональность

- **Размещение объявления** (`CreateAd`)
- **Получение ленты объявлений** с фильтрами и пагинацией (`ListAds`)
- Аутентификация пользователей через gRPC-запрос к `auth-service`
- Хранение данных в PostgreSQL


---

## gRPC API

### CreateAd

```proto
rpc CreateAd(CreateAdRequest) returns (AdResponse);

message CreateAdRequest {
  string title = 1;
  string description = 2;
  string image_url = 3;
  double price = 4;
  string token = 5;
}
```

- Требует авторизации через `auth-service.ValidateToken`.
- Возвращает данные добавленного объявления.

---

### ListAds

```proto
rpc ListAds(ListAdsRequest) returns (ListAdsResponse);

message ListAdsRequest {
  int32 page = 1;
  int32 page_size = 2;
  enum SortField { CREATED_AT = 0; PRICE = 1; }
  enum SortOrder { DESC = 0; ASC = 1; }
  SortField sort_by = 3;
  SortOrder sort_order = 4;
  double price_min = 5;
  double price_max = 6;
  string token = 7;
}
```

- Поддерживает пагинацию, фильтрацию по цене, сортировку.
- Для авторизованных пользователей выставляет `is_owner = true` для своих объявлений.

---

### Ответ `AdResponse`

```proto
message AdResponse {
  string id = 1;
  string title = 2;
  string description = 3;
  string image_url = 4;
  double price = 5;
  string author_login = 6;
  string created_at = 7;
  bool is_owner = 8;
}
```

---

## Конфигурация

Файл: `config/local.yaml`

```yaml
server:
  host: 0.0.0.0
  port: 50052

database:
  host: goods-db
  port: 5432
  user: goods_user
  password: goods_pass
  dbname: goods_db
  sslmode: disable

auth_service:
  address: auth-service:50051

redis:
  addr: redis:6379
  password: ""
  db: 0
```

---

## Авторизация

Все методы, кроме ListAds, требуют передачи JWT-токена. Проверка осуществляется через:

```
auth-service.ValidateToken(token)
```


---

## Архитектура

```
goods-service/
├── Makefile
├── README.md
├── api
│   ├── auth
│   │   └── auth.proto
│   ├── gen
│   │   ├── auth
│   │   │   ├── auth.pb.go
│   │   │   └── auth_grpc.pb.go
│   │   └── goods
│   │       ├── goods.pb.go
│   │       └── goods_grpc.pb.go
│   └── goods
│       └── goods.proto
├── cmd
│   └── main.go
├── config
│   └── local.yaml
├── deployments
│   └── Dockerfile
├── go.mod
├── go.sum
└── internal
    ├── config
    │   └── config.go
    ├── db
    │   ├── db.go
    │   ├── init_db.go
    │   ├── postgres.go
    │   └── schema.sql
    ├── grpc
    │   ├── auth_client.go
    │   ├── handler.go
    │   └── server.go
    ├── models
    │   └── models.go
    └── service
        ├── service.go
        └── service_impl.go
```