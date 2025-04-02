# Email-verification-api

## Стек технологий
- **Golang** (net/http)
- **SQL driver**
- **Docker**

## Установка и запуск

### 1. Клонирование репозитория
```sh
git clone https://github.com/enagtim/email-verification-api.git
cd  email-verification-api
```

### 2. Настройка переменных окружения
Создайте файл `.env` и добавьте в него:
```env
DSN=host=localhost user=user password=user dbname=email_verify port=5432 sslmode=disable
```

### 3. Запуск сервера
```sh
go run main.go
```

## Docker

### Запуск базы данных через Docker
```sh
docker-compose up -d
```

### Остановка контейнеров
```sh
docker-compose down
```