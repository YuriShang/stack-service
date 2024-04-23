# stack-service

## Команда для запуска
`docker compose up`

## **Конечные точки**
- `POST /push` - положить в стек
```json
{
    "data": 123
}
```
- `DELETE /pop` - достать из стека
```json
{
    "id": 1,
    "data": 123
}
```