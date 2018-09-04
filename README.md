# Queue

Server for help queue application

# Endpoints

- GET `/queue` lists all queues
- POST `/queue` adds a queue. Requires queue in body
- PATCH `/queue/{id}` toggles queue contacted status. Requires queue id
- DELETE `/queue/{id}` marks a queue as completed. Requires queue id

# Usage

- copy `.env.example` to `.env` and add config information for Mariadb or Mysql
- run migrations: `mysql -h DB_HOST -u DB_USERNAME -pDB_PASSWORD DB_NAME < ./queue_migrations.sql`

# Test

- create test db and run migrations as showned above
- run `go test`
