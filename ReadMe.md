# <div align="center">Тестовое задание Junior Golang Developer</div>
# <div align="center">Effective Mobile</div>



## 📝 Описание проекта

Микросервис для управления персональными данными с автоматическим обогащением информации через внешние API.

### 📋 Условия задачи

```text
✅ Реализовать CRUD API для работы с ФИО людей
✅ Интеграция с сервисами определения:
   • Возраст (Agify)
   • Пол (Genderize)
   • Национальность (Nationalize)
✅ Хранение данных в PostgreSQL
✅ Полная документация API через Swagger UI
✅ Поддержка фильтрации и пагинации
✅ Логирование операций
✅ Конфигурация через .env файл
```
____
## 🚀 Быстрый старт

Необходимые компоненты

• Go 1.20+

• PostgreSQL 14+

• Swag CLI

```bash
# Клонирование репозитория
git clone https://github.com/RakhimovAns/person-service.git
cd person-service

# Настройка окружения
cp .env.example .env
# Отредактируйте .env файл под вашу конфигурацию

# Установка зависимостей
go mod download

#Создание базы данных
psql -U postgres -c "CREATE DATABASE person_service;"

#Миграции
psql -U postgres -d person_service -f internal/repository/migrations/000001_create_people_table.up.sql

# Генерация документации Swagger
swag init -g cmd/main.go

# Запуск сервера
go run cmd/main.go
```

____
##  📚 Документация API

Доступна через Swagger UI после запуска сервера:

🔗 http://localhost:8080/swagger/index.html

____

## 📞 Контакты
<div align="center"> <a href="https://t.me/Rakhimov_Ans"> <img src="https://img.shields.io/badge/Telegram-Contact-blue?logo=telegram" alt="Telegram"> </a> <a href="https://github.com/RakhimovAns"> <img src="https://img.shields.io/badge/GitHub-Profile-black?logo=github" alt="GitHub"> </a> </div>
<div align="center"> <sub>Разработано с ❤️ для Effective Mobile</sub> </div> ```