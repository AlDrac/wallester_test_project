Wallister test project
=
A Golang v1.14 project created on September 11, 2020

### Start

- to run the project enter `make` into the console
```bash
make
```

### Migration
- enter the container to work with migrations
```bash
docker-compose exec -u 1000:1000 migration bash
```
or
```bash
make docker-bash-migration
```

###### Basic commands
- run migration - db url example 
`postgres://user:password@database.loc:5432/your_table?sslmode=disable` 
you can use environment variables `${POSTGRES_URL}` and `${POSTGRES_URL_TEST}`
```bash
migrate -path . -database {database_url} up
```
- create migration
```bash
migrate create -ext sql -dir . {migration_name}
```