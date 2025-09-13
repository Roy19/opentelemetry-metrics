## Opentelemetry Demo using Signoz

This demo showcases how to instrument applications using OpenTelemetry and visualize the traces in SigNoz.

First run migrations using go-migrate:

```bash
migrate -path db/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up
```

Then run initialize the database using some fake data:

```bash
chmod +x ./load_init_data.sh
./load_init_data.sh localhost 5432 user dbname
```

Run the application:

```
chmod +x ./start_the_application.sh
./start_the_application.sh
```