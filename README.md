# url_shortener

Команда для запуска контейнера:
```
docker run --env-file variables.env -d -p 8888:8888 url_shortener:1.0
```

В файле variables.env необходимо задать минимум 2 константы:
- Тип хранилища
- Путь к файлу с конфигом бд

Пример variables.env файла:

```
STORAGE="postgres"
ENV="C:\...\url_shortener\db.env"
```

Пример минимального db.env файла:
```
PASSWORD="123456"
```
