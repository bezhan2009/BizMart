
# BizMart

![Go Version](https://img.shields.io/badge/Go-1.22-blue)
![Swagger Version](https://img.shields.io/badge/Swagger-2.0-orange)

## Описание проекта
Этот проект представляет собой сервис для магазинов и пиццерий, которые хотят создать онлайн-присутствие, но не имеют бюджета на разработку собственного сайта. С помощью нашего сервиса пользователи смогут легко зарегистрировать свои магазины и начать продавать товары онлайн.

## Основные функции

### Часть с онлайн магазинами
- **Регистрация**: Пользователи могут зарегистрироваться на платформе.
- **Создание магазина**: После регистрации пользователи могут создать свой онлайн магазин, где им будут предоставлены функции для добавления и управления товарами.

### Часть с покупками
- **Покупатели**: Пользователи, которые не зарегистрировали свои магазины, могут выступать в роли покупателей.
- **Поиск товаров**: Главная страница позволит пользователям искать товары по названию. Вместо того чтобы анализировать множество интернет-магазинов, пользователи смогут ввести название товара и получить предложения от магазинов, зарегистрированных на платформе.

## Используемые технологии
- **API**: Для реализации функционала был использован GIN-GONIC.
- **Язык программирования**: Go 1.22
- **Документация**: Swagger 2.0

## Установка и запуск
1. Убедитесь, что у вас установлен Go 1.22.
2. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/BizMart.git
   ```
3. Перейдите в директорию проекта:
   ```bash
   cd BizMart
   ```
4. Установите зависимости:
   ```bash
   go mod tidy
   ```
5. Запустите проект:
   ```bash
   go run main.go
   ```

## Вклад
Если вы хотите внести свой вклад в проект, пожалуйста, создайте форк репозитория и отправьте пулл-реквест с вашими изменениями.

## Лицензия
Этот проект лицензирован под MIT License - подробности смотрите в файле [LICENSE](LICENSE).
