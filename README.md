# wap
Простой сервис с REST API с метриками, возвращающий текущее время по timezone

## Install
Для запуска вы можете использовать подготовленный docker контейнер из представленного Dockerfile,
указав через соответствующие переменные окружения параметры сервиса:
`docker run -d -p 8080:8080 --name wap -e LOGLEVEL=DEBUG alexdyukov/wap`

## REST API
Сервис поддерживает 2 endpoint:
* /metrics
* /time

И следующие методы:
* получить метрики: `GET /metrics`
* получить текущее время: `POST /time -d '{"timezone":"TIMEZONE_NAME"}'`
