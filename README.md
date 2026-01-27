# AutoOnline - CLI Automation Tool

![Docker](https://img.shields.io/badge/Docker-✓-blue?logo=docker)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
![License](https://img.shields.io/badge/License-MIT-green)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey)

Go-приложение для поддержания в онлайне на различных веб-сайтах, либо для фиксированной отправки запросов на сервер.

- [English Version README](#English-Guide)


##  Запуск программы

### Запуск через Docker
```bash
# сборка Docker образа
docker build -t autoonline .

# запуск в интерактивном режиме
docker run -it --rm autoonline
```

### Запуск через сборку
```bash
git clone https://github.com/aqvilov/AutoOnline.git
cd AutoOnline
go build -o autoonline main.go

# Запуск
./autoonline
#либо запуск через созданный .exe
```

##  Оглавление
- [Установка](#установка)
- [Использование](#использование)
- [Docker](#docker)
- [Тестирование](#Unit-тесты)


## 🐳 Docker

### Быстрый старт с Docker
```bash
# 1. Клонируйте репозиторий
git clone https://github.com/aqvilov/AutoOnline.git
cd AutoOnline

# 2. Соберите Docker образ
docker build -t autoonline .

# 3. Запустите контейнер
docker run -it --rm autoonline
```

### Использование Docker Compose
Файл `docker-compose.yml` позволяет запускать приложение одной командой:

```bash
# Запуск в интерактивном режиме
docker-compose up

# Запуск в фоновом режиме
docker-compose up -d

# Просмотр логов
docker-compose logs -f

# Остановка
docker-compose down
```


## 🛠 Установка (без Docker)

### Требования
- Go 1.25
- Git (для клонирования репозитория)

### Шаги установки
```bash
# 1. Клонирование репозитория
git clone https://github.com/aqvilov/AutoOnline.git
cd AutoOnline

# 2. Установка зависимостей
go mod download

# 3. Сборка приложения
go build -o autoonline main.go

# 4. Проверка установки
./autoonline --help
```

### Удобный запуск на Windows (альтернативный способ через run.bat)
```batch
:: Используйте run.bat для запуска
run.bat
```

## 📖 Использование

### Запуск приложения
```bash
# После установки
./autoonline
```

### Процесс работы
1. **Ввод URL**: Введите корректный URL сайта
2. **Настройка интервала**: Укажите интервал между запросами (в минутах)
3. **Время работы**: Задайте общее время работы или 0 для бесконечного режима
4. **Cookie аутентификация**: Следуйте инструкциям для получения cookies
5. **Запуск**: Программа начнет отправлять запросы

### Инструкция по получению Cookies ( написано в программе )
При запуске программа покажет подробную инструкцию:
1. Откройте инструменты разработчика (F12)
2. Перейдите на вкладку Application/Storage → Cookies
3. Найдите домен вашего сайта
4. Скопируйте все пары "Name=Value"
5. Вставьте в формате: `name1=value1; name2=value2; name3=value3`

## Unit-тесты

### Запуск тестов
```bash
# Нативно
go test ./... -v

# Через Docker
docker build -t autoonline-test .
docker run --rm autoonline-test go test ./...
```


**aqvilov**
- GitHub: [@aqvilov](https://github.com/aqvilov)

---
Приложение запрашивает Cookie-файлы.
Программа полностью безопасна с открытым исходным кодом и не хранит никакие ваши данные
---




##  English Guide

Go application for staying online on various websites or for sending fixed requests to the server.

##  Running the program

### Running via Docker
```bash
# building the Docker image
docker build -t autoonline .

# run in interactive mode
docker run -it --rm autoonline
```

### Running via build
```bash
git clone https://github.com/aqvilov/AutoOnline.git
cd AutoOnline
go build -o autoonline main.go

# Run
./autoonline
#or run via the created .exe
```

##  Table of contents
- [Installation](#installation)
- [Usage](#Using-Docker-Compose)
- [Docker](#docker)
- [Testing](#Unit-tests)


## 🐳 docker

### Quick start with Docker
```bash
# 1. Clone the repository
git clone https://github.com/aqvilov/AutoOnline.git
cd AutoOnline

# 2. Build the Docker image
docker build -t autoonline .

# 3. Run the container
docker run -it --rm autoonline
```

### Using Docker Compose
The `docker-compose.yml` file allows you to run the application with a single command:

```bash
# Run in interactive mode
docker-compose up

# Run in the background
docker-compose up -d

# View logs
docker-compose logs -f

# Stop
docker-compose down
```


## 🛠 Installation (without Docker)

### Requirements
- Go 1.25
- Git (for cloning the repository)

### Installation steps
```bash
# 1. Clone the repository
git clone https://github.com/aqvilov/AutoOnline.git
cd AutoOnline

# 2. Install dependencies
go mod download

# 3. Build the application
go build -o au

# 4. Checking the installation
./autoonline --help
```

### Convenient launch on Windows (alternative method via run.bat)
```batch
:: Use run.bat to launch
run.bat
```

## 📖 Usage

### Launching the application
```bash
# After installation
./autoonline
```

### Workflow
1. **Enter URL**: Enter the correct website URL
2. **Set interval**: Specify the interval between requests (in minutes)
3. **Run time**: Set the total run time or 0 for infinite mode
4. **Cookie authentication**: Follow the instructions to obtain cookies
5. **Start**: The program will start sending requests

### Instructions for obtaining cookies (written in the program)
When launched, the program will display detailed instructions:
1. Open the developer tools (F12)
2. Go to the Application/Storage → Cookies tab
3. Find your website's domain
4. Copy all “Name=Value” pairs
5. Paste them in the following format: `name1=value1; name2=value2; name3=value3`

## Unit tests

### Running tests
```bash
# Native
go test ./... -v

# Via Docker
docker build -t autoonline-test .
docker run --rm autoonline-test go test ./...
```


**aqvilov**
- GitHub: [@aqvilov](https://github.com/aqvilov)

---


---
