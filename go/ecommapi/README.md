# ecommapi
**EcommAPI** is an implementation of REST APIs for e-commerce platforms.

## Tech Stack
- Go
- PostgreSQL
- GORM
- Gin Web Framework

## Installation

1. Clone the repository:
  ```sh
  git clone <repository-url>
  ```
2. Create a `.env` file, set the database connection and update the port and postgres credentials with the actual values:
  ```sh
  PORT=8081
  dsn="host=localhost user=postgresuser password=postgrespass dbname=mydb port=5432 sslmode=disable"
  ```

4. Setup the database:
  ```sh
  docker compose up
  ```

6. Then start the API:
  ```sh
  go run ./cmd/ecommapi
  ```

7. (Optional) Run the tests in `tests` folder:
  ```sh
  go test . -v
  ```

8. (Optional) Generate updated docs:
  ```sh
  swag init
  OR
  swag init -g cmd/ecommapi/main.go -o docs
  ```


## Functionalities
- Authentication: Stateful Sessions (token-based with database validation)
- Header check: HTTP header X-Session-Token for non-public requests
- Password hashing with bcrypt
  - https://pkg.go.dev/golang.org/x/crypto/bcrypt
  - https://pkg.go.dev/github.com/absagar/go-bcrypt#section-readme
- UUID for keys https://github.com/lithammer/shortuuid
- Sessions data in a database table
  - Adding the token upon logging in
  - Removing the token upon logging out
  - Middleware to check admin status
- No JSONB
- Product Catalogs
  - Pagination (10 products per page)
- Test case for each endpoint with more than 90% code coverage
- Documentation using Swagger UI (swaggo)


## Database Schema

#### Users
```
  id: uuid
  username: varchar
  email: varchar
  password_hash: text
  role: varchar
```


#### Sessions
```
  id: uuid
  user_id: uuid
  token: varchar
  expires_at: timestamp
```

#### Orders
```
  id: uuid
  user_id: uuid
  status: varchar
  total_amount: decimal
  shipping_street: varchar
  shipping_city: varchar
  shipping_zip: varchar
  shipping_country: varchar
```

#### Products
```
  id: uuid
  name: varchar
  description: text
  price: decimal
  category_id: varchar
  is_active: bool
```

#### Stocks
```
  product_id: uuid
  quantity: integer
  reserved: bool
  location: varchar
```

#### Order_Items
```
  order_id: uuid
  product_id: uuid
  quantity: integer
  unit_price: decimal
```

#### Cart_Items
```
  cart_id: uuid
  product_id: uuid
  quantity: integer
```

#### Notifications
```
  id: uuid
  user_id: uuid
  title: varchar
  message: text
  is_read: bool
```


## API Usage

#### POST /auth/register
- **Request Body:**
  ```json
  {
    "username": "string",
    "email": "string",
    "password": "string"
  }
  ```
- **Response Example (201 Created):**
  ```json
  {
    "id": "uuid",
    "username": "string",
    "email": "string"
  }
  ```

#### POST /auth/login
- **Request Body:**
  ```json
  {
    "email": "string",
    "password": "string"
  }
  ```
- **Response Example (200 OK):**
  ```json
  {
    "sesion_token": "string",
    "expires_at": "timestamp",
    "user": {
      "username": "string",
      "role": "Customer / Admin"
    }
  }
  ```

#### POST /auth/logout
- **Request:**
  ```
  X-Session-Token (Customer / Admin)
  ```
- **Response Example (204 No Content):**
  ```json
  {
    "sesion_token": "string",
    "expires_at": "timestamp",
    "user": {
      "username": "string",
      "role": "Customer / Admin"
    }
  }
  ```

#### GET /products
- **Query Parameters:**
  ```
  category_id: "string"
  page: int
  ```
- **Response Example (200 OK):**
  ```json
  {
    "items": [
      {
        "name": "string",
        "price": decimal,
        "category": "string"
      },
      ...
    ],
    "total_count": decimal
  }
  ```

#### POST /products
- **Request:**
  ```
  X-Session-Token (Admin)
  ```
- **Request Body:**
  ```json
  {
    "name": "string",
    "price": decimal,
    "description": "string",
    "category_id": "string"
  }
  ```
- **Response Example (201 Created):**
  ```json
  {
    "id": "uuid",
    "name": "string",
    "price": decimal,
    "description": "string",
    "category_id": "string",
    "is_active": bool
  }
  ```

#### GET /inventory/:product_id
- **Request:**
  ```
  X-Session-Token (Customer / Admin)
  ```
- **Response Example (200 OK):**
  ```json
  {
    "quantity": integer,
    "stock": integer,
    "storageLocation": "string"
  }
  ```


#### PATCH /inventory/:product_id
- **Request:**
  ```
  X-Session-Token (Admin)
  ```
- **Request Body:**
  ```json
  {
    "increment_by": integer,
    "reason": "string"
  }
  ```
- **Response Example (200 OK):**
  ```json
  {
    "quantity": integer,
    "stock": integer,
    "storageLocation": "string"
  }
  ```


#### GET /cart
- **Request:**
  ```
  X-Session-Token (Customer)
  ```
- **Response Example (200 OK):**
  ```json
  {
    "cart_items": [
      {
        "name": "string",
        "price": decimal,
        "category": "string"
      },
      ...
    ],
    "total_amount": decimal
  }
  ```


#### POST /cart/items
- **Request:**
  ```
  X-Session-Token (Customer)
  ```
- **Request Body:**
  ```json
  {
    "product_id": uuid,
    "quantity": integer
  }
  ```
- **Response Example (201 Created):**
  ```json
  {
    "items": [
      {
        "name": "string",
        "price": decimal,
        "category": "string"
      },
      ...
    ],
    "total_amount": decimal
  }
  ```

#### POST /orders
- **Request:**
  ```
  X-Session-Token (Customer)
  ```
- **Request Body:**
  ```json
  {
    "items": [
      {
        "product_id": uuid,
        "quantity": integer
      },
      ...
    ],
    "shipping_street": "string",
    "shipping_city": "string",
    "shipping_zip": "string",
    "shipping_country": "string"
  }
  ```
- **Response Example (201 Created):**
  ```json
  {
    "items": [
      {
        "name": "string",
        "price": decimal,
        "category": "string"
      },
      ...
    ],
    "shipping_street": "string",
    "shipping_city": "string",
    "shipping_zip": "string",
    "shipping_country": "string",
    "total_amount": decimal
  }
  ```


#### GET /notifications
- **Request:**
  ```
  X-Session-Token (Customer)
  ```
- **Response Example (200 OK):**
  ```json
  {
    "notifications": [
      {
        "username": "string",
        "title": "string",
        "message": "string",
        "is_read": bool
      },
      ...
    ]
  }
  ```

#### PATCH /notifications/:id/read
- **Request:**
  ```
  X-Session-Token (Customer)
  ```
- **Response Example (200 OK):**
  ```json
  {
    "username": "string",
    "title": "string",
    "message": "string",
    "is_read": true
  }
  ```

#### Other Status Codes
- 200 OK
- 201 Created
- 400 Bad Request
- 401 Unauthorized
- 403 Forbidden
- 404 Not Found
- 500 Internal Server Error


## Helper Commands

- Backup DB Table data
```
docker exec -t ecommapi-db-1 psql -U postgresuser -d mydb -c "\copy products TO STDOUT WITH CSV HEADER" > db.csv
```

- Import DB Table data
```
docker exec -i ecommapi-db-1 psql -U postgresuser -d mydb -c "\copy products FROM STDIN WITH CSV HEADER" < db.csv
```