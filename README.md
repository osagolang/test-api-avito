# Test API Avito

Тестовый проект - внутренняя система учета пользователей, кошельков, транзакций и покупок товаров.

## Описание

Проект реализует REST API для внутренней системы, в которой пользователи могут:

- регистрироваться и авторизовываться с помощью JWT
- просматривать баланс кошелька
- совершать транзакции между пользователями
- покупать товары в магазине
- просматривать инвентарь

Технологии:
- Go (Gin)
- PostgreSQL
- Docker + Docker Compose
- JWT (github.com/golang-jwt/jwt/v5)
- bcrypt (golang.org/x/crypto/bcrypt)
- SQL-миграции (github.com/rubenv/sql-migrate)

---

How to install
```bash
docker-compose up -d
go install github.com/rubenv/sql-migrate/...@latest
sql-migrate up
```
