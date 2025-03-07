# Тестовое задание SkillsRock
## **Инструкция по запуску**
### Клонируйте репозиторий:

```bash
git clone https://github.com/zatrasz75/SkillsRock.git
```
* Перейдите в директорию проекта:

```bash
cd SkillsRock
```
* Запуск проекта с помощью Docker Compose:
```bash
make dock
```
Запуск сервера на http://localhost:3000

Документация Swagger API: http://localhost:3000/swagger/index.html

### Запустить локально
* Отредактируйте файл .env 
```bash
APP_IP=localhost
APP_PORT=3000
CORS_ALLOWED_ORIGINS=localhost

# Для миграции бд
DB_CONNECTION_STRING=postgres://zatrasz:postgrespw@postgres_rock:5432/db_rock?sslmode=disable

# Не обязательно если сипользуется DB_CONNECTION_STRING
POSTGRES_USER=zatrasz
POSTGRES_PASSWORD=postgrespw
HOST_DB=postgres_rock
PORT_DB=5432
POSTGRES_DB=db_rock
```
* Запуск приложения
```bash
make run
```
Запуск сервера на http://localhost:3000

Документация Swagger API: http://localhost:3000/swagger/index.html

### endpoints:
✅ POST /tasks – создание задачи.

✅ GET /tasks – получение списка всех задач.

✅ PUT /tasks/:id – обновление задачи.

✅ DELETE /tasks/:id – удаление задачи.

### Создание миграции для БД
* создает файл с уникальным именем из даты и времени в /migrations
```bash
make up
```
* Откатить в базе миграцию
```bash
make down
```
### Логирование
* Записывает отформатированное сообщение лога в консоль и в файл ./var/log/main.log"
#### Уровни логирования
- Info (зеленый)

- Success (синий)

- Trace (серый)

- Error (красный)

- Warn (желтый)

- Fatal (красный)

- Debug (голубой)

- Critical (фиолетовый)

- Panic (оранжевый)

- Security (яркий фиолетовый)