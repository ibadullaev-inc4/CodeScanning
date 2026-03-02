# TODO Service Microservice - Clean Architecture

Микросервис на Go с чистой архитектурой (Clean Architecture / Hexagonal Architecture).

## 🚀 Быстрый старт

### Требования разработчика
```bash
# Установить pre-commit (один раз)
pip install pre-commit

# Активировать правила в репозитории (после клонирования)
pre-commit install
```

Подробнее: [CONTRIBUTING.md](./CONTRIBUTING.md)

## Структура проекта

```
.
├── main.go                          # Точка входа, инициализация зависимостей
├── go.mod                           # Модуль Go
└── internal/
    ├── domain/                      # Domain layer - сущности
    │   └── todo.go                  # Модель Todo (no dependencies)
    ├── port/                        # Ports layer - интерфейсы (абстракции)
    │   └── port.go                  # TodoRepository, TodoService интерфейсы
    ├── service/                     # Application/UseCase layer
    │   └── todo.go                  # Бизнес-логика (service)
    └── adapter/                     # Adapter layer - реализации
        ├── http/                    # HTTP адаптер
        │   └── handler.go           # HTTP контроллеры
        └── storage/                 # Storage адаптер
            └── memory.go            # In-memory реализация репозитория
```

## Слои архитектуры

1. **Domain Layer** (`internal/domain/`)
   - Чистые бизнес-сущности без зависимостей
   - Не знает о базах данных, HTTP, фреймворках

2. **Port Layer** (`internal/port/`)
   - Интерфейсы (порты)
   - Определяют контракты между слоями
   - Позволяют менять реализации без изменения бизнес-логики

3. **Service/UseCase Layer** (`internal/service/`)
   - Реализация бизнес-логики
   - Зависит только от интерфейсов (Port)
   - Не зависит от конкретных реализаций

4. **Adapter Layer** (`internal/adapter/`)
   - HTTP: контроллеры для web запросов
   - Storage: реализации хранилищ (памяти, БД, кэша и т.д.)
   - Могут быть заменены без изменения ядра

## API Endpoints

- `GET /health` — проверка статуса
- `GET /todos` — получить все задачи
- `POST /todos` — создать новую задачу
- `GET /todos/{id}` — получить задачу по ID
- `PUT /todos/{id}` — обновить задачу
- `DELETE /todos/{id}` — удалить задачу

## Запуск

```bash
go mod download
go run main.go
```

Сервис будет доступен на `http://localhost:8080`

## Преимущества архитектуры

✅ **Модульность** — легко добавлять новые функции
✅ **Тестируемость** — можно мокировать интерфейсы
✅ **Гибкость** — просто менять реализации (SQL БД вместо памяти)
✅ **Масштабируемость** — чистое разделение ответственности
✅ **Независимость** — фреймворки не диктуют архитектуру
