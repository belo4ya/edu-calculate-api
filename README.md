# Edu Calculate API

Программирование на Go | 24. Веб-сервис для вычисления арифметических выражений.
Пользователь отправляет арифметическое выражение по HTTP и получает в ответ его результат.

## 🚀 Запуск

Запустить:

```shell
make run
```

Запустить линтер:

```shell
make lint
```

Запустить тесты с отчетом о покрытии:

```shell
make test-cover
```

## 💡 Использование

У сервиса 1 endpoint с url-ом `/api/v1/calculate`. Пользователь отправляет на этот url POST-запрос с телом:

```shell
curl --write-out 'status code: %{http_code}\n' localhost:8080/api/v1/calculate \
-H 'Content-Type: application/json' \
--data '{"expression": "2+2*2"}'
```

Ответ:

```text
{"result":6}
status code: 200
```

Запрос с невалидным выражением:

```shell
curl --write-out 'status code: %{http_code}\n' localhost:8080/api/v1/calculate \
-H 'Content-Type: application/json' \
--data '{"expression": "2+2*"}'
```

Ответ:

```text
{"error":"Expression is not valid"}
status code: 422
```

Невалидный запрос:

```shell
curl --write-out 'status code: %{http_code}\n' localhost:8080/api/v1/calculate \
-H 'Content-Type: application/json' \
--data '{"expression": 10}'
```

Ответ:

```text
{"error":"Internal server error"}
status code: 500
```

## 📚 Структура

```text
├── Dockerfile
├── Makefile
├── README.md
├── bin
│   └── server
├── cmd
│   ├── cli
│   │   └── main.go
│   └── server
│       └── main.go
├── coverage
│   ├── cover
│   └── cover.html
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── httputil
│   │   └── httputil.go
│   ├── logging
│   │   └── logging.go
│   ├── server
│   │   └── http.go
│   └── service
│       ├── calculate.go
│       ├── calculate_test.go
│       └── service.go
└── pkg
    └── calc
        ├── calc.go
        ├── calc_test.go
        └── errors.go

```
