Wallester test project
=
A Golang v1.14 project created on September 11, 2020

### Start

- to run the project enter `make` into the console
```bash
make
```

after 10 second you can use websites

- API urls
```
[ GET ]     - http://localhost/api/v1/customers
// JSON REQUEST BODY {"first_name": "", "last_name": ""}

[ GET ]     - http://localhost/api/v1/customer/{id:[0-9]+}
[ POST ]    - http://localhost/api/v1/customer/create
// JSON REQUEST BODY
{
	"first_name": "Test40",
	"last_name": "Tester40",
	"birth_date": "2001-05-11",
	"gender": "male",
	"email": "test40.tester@test.com",
	"password": "asdasd@543!asasd",
	"address": "Test address 40"
}

[ PUT ]     - http://localhost/api/v1/customer/edit/{id:[0-9]+}
// JSON REQUEST BODY
{
	"first_name": "Test50",
	"last_name": "Tester50",
	"birth_date": "2001-05-11",
	"gender": "male",
	"email": "test40.tester@test.com",
	"address": "Test address 40"
}

[ DELETE ]  - http://localhost/api/v1/customer/delete/{id:[0-9]+}
```

### Close
- to close the project enter `make docker-down` into the console
```bash
make docker-down
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
