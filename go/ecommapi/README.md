# ecommapi
**EcommAPI** is an implementation of REST APIs for e-commerce platforms.

## Tech Stack
- Go
- PostgreSQL
- GORM

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
- Product Catalogs
  - Pagination (10 products per page)
- No JSONB
- Test case for each endpoint with more than 90% code coverage
- Documentation using Swagger UI (swaggo)


## Database Schema

#### Users
```json
  id: uuid
  username: varchar
  email: varchar
  password_hash: text
  role: varchar
```


#### Sessions
```json
  id: uuid
  user_id: uuid
  token: varchar
  expires_at: timestamp
```

#### Orders
```json
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
```json
  id: uuid
  name: varchar
  description: text
  price: decimal
  category_id: varchar
  is_active: bool
```

#### Stocks
```json
  product_id: uuid
  quantity: integer
  reserved: bool
  location: varchar
```

#### Order_Items
```json
  order_id: uuid
  product_id: uuid
  quantity: integer
  unit_price: decimal
```

#### Cart_Items
```json
  cart_id: uuid
  product_id: uuid
  quantity: integer
```

#### Notifications
```json
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
  ```json
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
  ```json
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
  ```json
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

#### GET /inventory/:product_id
- **Request:**
  ```json
  X-Session-Token (Customer / Admin)
  ```
- **Response Example (200 OK):**
  ```json
  {
    "quantity": integer,
    "stock": integer,
    "string": storageLocation
  }
  ```


#### PATCH /inventory/:product_id
- **Request:**
  ```json
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
    "string": storageLocation
  }
  ```


#### GET /cart
- **Request:**
  ```json
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
  ```json
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
  ```json
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
  ```json
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
  ```json
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