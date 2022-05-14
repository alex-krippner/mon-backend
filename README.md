# Mon backend

---

## Table of Contents

1. [General Info](#general-info)
2. [Installation and build](#installation-and-build)
3. [Database migration](#database-migration)

### General Info

---

Dockerized backend for mon written in Go with a Postgres database.

## Installation and Build

---

Build container  
`docker compose build`

Init Go module  
`docker compose run --rm app go mod init mon-backend`

Setup Air  
`docker compose run --rm app air init`

Run  
`docker compose up`

Run with pgadmin
`docker compose --profile pgadmin up`

Update modules  
`docker compose run --rm app go mod tidy`

## Database migration

---

Create new migrations  
`docker compose --profile tools run create-migration <migration_filename>`

Run migration  
`docker compose --profile tools run migrate`

Down migration  
`docker compose --profile tools run migrate-down`
