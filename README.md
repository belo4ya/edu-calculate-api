# Edu Calculate API

Программирование на Go | 24. Веб-сервис для вычисления арифметических выражений.
Пользователь отправляет арифметическое выражение по HTTP и получает в ответ его результат.

## 🚀 Запуск

Запустить сервер:

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
├── cmd
│   ├── cli
│   │   └── main.go  # точка входа для запуска приложения в режиме командной строки
│   └── server
│       └── main.go  # точка входа для запуска приложения в режиме веб-сервиса
├── internal
│   ├── config
│   │   └── config.go  # структура хранения и загрузки конфига
│   ├── httputil
│   │   └── httputil.go
│   ├── logging
│   │   └── logging.go  # конфигурация логгера
│   ├── server
│   │   └── http.go  # конфигурация HTTP-сервера
│   └── service
│       ├── calculate.go  # бизнес-логика обработчика запросов
│       ├── calculate_test.go
│       └── service.go
└── pkg
│   └── calc  # модуль вычисления арифметических выражений
│       ├── calc.go
│       ├── calc_test.go
│       └── errors.go
├── Makefile  # полезные команды
└── README.md  # вы здесь
```
