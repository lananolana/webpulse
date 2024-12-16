# WebPulse
Платформа для мониторинга состояния и производительности веб-сайтов.

# Quick start

Склонировать репозиторий
```bash
git clone https://github.com/lananolana/webpulse.git
```

Создать конфиг приложения в папке infra

```bash
cd webpulse/infra && nano app.yaml
```

`app.yaml`
```yaml
app:
  # mock [true, false] - Включает моковые ответы на запросы
  mock: false

  # log_level: [DEBUG, INFO, WARNING, ERROR]
  log_level: INFO

  # log_format: [text, json]
  log_format: json

  http_server:
    listen_addr: 0.0.0.0:8080
    read_timeout: 10s
    write_timeout: 10s
    idle_timeout: 10s

  http_client:
    timeout: 10s
```

Собрать и запустить контейнер с webpulse backend
```bash
docker-compose -f infra/docker-compose.yml up -d --build
```

Swagger:
```
http://localhost:8080/docs/index.html
```
