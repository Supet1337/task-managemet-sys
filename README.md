### task-managemet-sys
Задание от chatGPT:
Вот мое предложение:

Создайте систему управления задачами (Task Management System) с использованием языка Go, микросервисной архитектуры, RabbitMQ, Docker, PostgreSQL и REST API.

Система должна состоять из трех микросервисов:
1. Сервис для создания и управления задачами (Task Service).
2. Сервис для управления пользователями (User Service).
3. Сервис для отправки уведомлений (Notification Service).

Каждый сервис должен быть запакован в Docker-контейнер и развернут в своем собственном контейнере.

Task Service должен иметь следующие функции:
- Создание новых задач.
- Получение списка задач.
- Обновление задач.
- Удаление задач.
- Установка статуса задачи (выполнена/не выполнена).
- Отправка уведомления о новой задаче.

User Service должен иметь следующие функции:
- Создание новых пользователей.
- Получение списка пользователей.
- Аутентификация пользователей.

Notification Service должен иметь следующую функцию:
- Отправка уведомления на почту при создании новой задачи.

Все сервисы должны взаимодействовать между собой через RabbitMQ. Task Service должен отправлять сообщения в Notification Service при создании новой задачи.

Для хранения данных используйте PostgreSQL.

REST API должен быть реализован для каждого сервиса. Используйте Swagger для документирования API.

В конечном итоге вы получите полноценную систему управления задачами, которая покрывает все основные аспекты микросервисной архитектуры и использует самые популярные технологии для разработки веб-приложений на Go.
