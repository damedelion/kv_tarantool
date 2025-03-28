# kv_tarantool

key-value хранилище с хранением данных в Tarantool

API:
- POST /kv body: {key: "test", "value": {SOME ARBITRARY JSON}} 
- PUT kv/{id} body: {"value": {SOME ARBITRARY JSON}}
- GET kv/{id} 
- DELETE kv/{id}

Ответы:
- POST  возвращает 409 если ключ уже существует, 
- POST, PUT возвращают 400 если боди некорректное
- PUT, GET, DELETE возвращает 404 если такого ключа нет

Запуск:
```
docker compose build
docker compose up
```

Завершение:
```
docker compose down
```