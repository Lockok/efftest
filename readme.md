# Effective Mobile — Test Task

REST API для управления подписками пользователей.

## Технологии

- Go
- PostgreSQL
- pgx
- Goose
- Docker / Docker Compose
- Swagger (OpenAPI)
- Kubernetes
- Helm
- GitHub Actions

## Архитектура

Проект построен с использованием принципов Clean Architecture.

```
Handler
    ↓
Service
    ↓
Repository
    ↓
PostgreSQL
```

Структура проекта:

```
cmd/
docs/
internal/
    config/
    domain/
    dto/
    handler/
    middleware/
    repository/
    service/
    server/
    storage/
migrations/
chart/
k8s/
```

## Возможности

- Создание подписки
- Получение подписки по ID
- Получение списка подписок пользователя
- Частичное обновление подписки (`PATCH`)
- Удаление подписки
- Расчёт суммарной стоимости подписок за выбранный период
- Фильтрация расчёта по пользователю и названию подписки

## Структура подписки

Подписка содержит:

- название
- стоимость
- UUID пользователя
- дату начала
- необязательную дату окончания

## Запуск проекта

### 1. Клонировать репозиторий

```bash
git clone https://github.com/Lockok/efftest.git
cd efftest
```

### 2. Создать файл `.env`

Linux/macOS:

```bash
cp .env.example .env
```

Windows:

```cmd
copy .env.example .env
```

При необходимости измените значения переменных окружения.

### 3. Запустить PostgreSQL

```bash
docker compose up -d
```

### 4. Применить миграции

```bash
make migrate-up
```

### 5. Запустить приложение

```bash
make app-run
```

После запуска API будет доступно по адресу:

```
http://localhost:8080
```

## Swagger

Документация:

```
http://localhost:8080/swagger/index.html
```

## Особенности

- `end_date` является необязательным полем.
- Даты передаются в формате `YYYY-MM`.
- Для обновления используется `PATCH`, поэтому можно передавать только изменяемые поля.
- Реализовано логирование HTTP-запросов.
- Реализован Recovery middleware для обработки panic.
- Каждому запросу автоматически присваивается `X-Request-ID`.

## Дополнительно

Проект также содержит:

- Docker Compose
- Kubernetes manifests
- Helm Chart
- GitHub Actions (CI)
- golangci-lint
