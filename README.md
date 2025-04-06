# TODO-лист на Go + Fiber + PostgreSQL

Простой сервис управления задачами (TODO-лист), написанный на языке Go с использованием веб-фреймворка [Fiber](https://github.com/gofiber/fiber) и базой данных PostgreSQL, развернутой в Docker.

## Стек технологий

- Go 1.24
- Fiber v5
- PostgreSQL (в контейнере Docker)
- YAML для конфигурации
- Поддержка `.env` для задания пути к конфигу

---

## Запуск приложения

### 1. Клонировать репозиторий

```bash
git clone https://github.com/legenda-hortici/todo-service.git
cd todo-service
```

### 2. Настроить .env файл
Создайте файл .env в корне проекта:
```env
CONFIG_PATH=config/config.yaml
```

### 3. Проверить конфигурацию

### 4. Запустить контейнер с базой данных PostgreSQL

```bash
docker run --name postgres_alt \
  -e POSTGRES_PASSWORD=0897 \
  -e POSTGRES_DB=todolist \
  -p 5433:5432 \
  -d postgres
```

### 5. Запустить приложение
```bash
go run cmd/api/main.go
```
---

Приложение запускается по адресу: http://localhost:8080

Можно тестировать с помощью:

- Postman

- curl

##  Возможности
- Добавление задач

- Обновление статуса

- Удаление задач

- Получение списка задач

- Статусы: new, in_progress, done
