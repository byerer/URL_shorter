# URL_shorter

## todo
- pgcode
- pkg/errors
- 读写分离
- 分库分表

## postgres
```
docker run --name url_shorter `
-e POSTGRES_PASSWORD=gorm `
-e POSTGRES_USER=gorm `
-e POSTGRES_DB=shorter `
-d -v url_shorter_postgres_data:/var/lib/postgresql/data `
-p 5432:5432 postgres
```