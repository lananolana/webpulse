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

Собрать и запустить контейнер с webpulse_backend и nginx
```bash
docker-compose -f infra/docker-compose.yml up -d --build
```

Сервер будет доступен на localhost:80, на котором висит nginx, проксирующий
запросы по сети docker в webpulse_backend (недоступен извне)


## Swagger
На странице со swagger есть два сервера на выбор:
- http://localhost/docs/index.html - работает при запуске через docker-compose (наиболее предпочтительный способ)
- http://localhost:8080/docs/index.html - работает, если запускаться без docker-compose и без nginx (для локальной разработки бекенда и тестов)

