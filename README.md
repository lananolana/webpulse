# WebPulse
Платформа для мониторинга состояния и производительности веб-сайтов.

# Quick start

Склонировать репозиторий
```bash
git clone https://github.com/lananolana/webpulse.git
```

Создать конфиг приложения

```bash
cd webpulse/infra && nano app.yaml
```

`app.yaml`
```yaml
app:
  # mock [true, false]
  mock: true

    # log_level: [DEBUG, INFO, WARNING, ERROR]
  log_level: INFO

  # log_format: [text, json]
  log_format: json

  http:
    listen_addr: 0.0.0.0:8080
```

Собрать и запустить контейнер с webpulse backend
```bash
docker-compose -f infra/docker-compose.yml up -d --build
```

