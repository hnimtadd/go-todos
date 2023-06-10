# TODOs app with clean architecture

## just for study purpose
### TechStack:
- Architecture: ```Clean architecture```
- ORM: ```Gorm```
- HTPP framework: ```Echo```
- DB: ```Postgres```
- Deployment: ```Docker```
- DB dashboard: ```Adminer```

### Overview
- Use JWT for validation
- Limit todo per user in a day

### What nexts ?
- Add some necessary unit tests and interfration tests
- Implement new features base on user needed
- Add role/permission based validation

### How to run:
- Clone project and change directory to project root's folder
- update your .env file with this format:
- Run locally:
    ```txt
    PORT=8080
    JWT_SECRET=B5bJHoI8aVLjAAeV
    SIGNING_KEY=ABC
    HASH_SALT=SJSHDFDS
    TOKEN_TTL=86400
    CONNECTION_URL=host=localhost user=postgres password=password1 dbname=todos port=5432
    ```
    then
    ```bash
    go run cmd/api/main.go
    ```
- Run by Docker:
    ```txt
    #...
    CONNECTION_URL=host=postgres user=postgres password=password1 dbname=todos port=5432
    ```
    then
    ```bash
    docker-compose up -d
    ```
- Manager postgres state by adminer: access to ```https:localhost:8060```
### api:

#### Register:

```bash
curl --location --request POST 'http://localhost:8080/api/v1/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "yourusername",
    "password": "yourpassword",
    "limit": 2
}'
```

#### Login:

```bash
curl --location --request POST 'http://localhost:8080/api/v1/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "yourusername",
    "password": "yourpassword"
}'
```

Take the token from login then.

#### Add todo:

```bash
curl --location --request POST 'http://localhost:8080/api/v1/todos/create' \
--header 'Authorization: Bearer $(token)' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content": "todo content"
}'
```

#### Get all todos:

Public for all users

```bash
curl --location --request GET 'http://localhost:8080/api/v1/todos/'
```

#### Get user todos:


```bash
curl --location --request GET 'http://localhost:8080/api/v1/todos/user' \
--header 'Authorization: Bearer YOUR_TOKEN_HERE'
```

