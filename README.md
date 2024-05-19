# URL_shorter

## todo
- pkg/errors
- 读写分离
- 分库分表

## postgres source
```
docker run --name url_shorter `
-e POSTGRES_PASSWORD=gorm `
-e POSTGRES_USER=gorm `
-e POSTGRES_DB=shorter `
-d -v url_shorter_postgres_data:/var/lib/postgresql/data `
-p 5432:5432 postgres
```



## redis
```
docker run --name url_shorter_redis `
-d -v url_shorter_redis_data:/data `
-p 6379:6379 redis
```