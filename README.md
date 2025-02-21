## Structure
```
/fitness-api
│── /cmd
│   ├── /handlers        # HTTP request handlers (controllers)
│   ├── /repositories   # Database access logic (queries)
│   ├── /services       # Business logic (service layer)
│   ├── /models         # Data models (structs)
│   ├── /storage        # Database connection setup
│── main.go             # Entry point of the application
```

- handlers/ → Handles API requests.
- repositories/ → Contains functions to fetch and modify data in the database.
- services/ → Contains business logic.
- models/ → Defines data structures.
- storage/ → Handles database connection.

## create table users 
```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
); 
```

#### set default value for datbase table

```sql
alter table users alter column created_at set default now();
```
```sql
alter table users alter column created_at set default now();
```

## create measurement table
```sql
CREATE TABLE measurements (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  weight FLOAT NOT NULL,
  height FLOAT NOT NULL,
  body_fat FLOAT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);
```
#### set default value for measurements table
```sql
alter table measurements alter column created_at set default now();
```