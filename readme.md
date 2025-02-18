# E-commerce Backend (Go + Gin + GORM)

This is a fully functional **E-commerce Backend** built using **Go**, **Gin**, and **GORM**. It includes authentication, product management, order management, and a review & ratings system.

![swagger](/swagger.png)

## 🚀 Features
- **User Authentication** (JWT-based)
- **Product Management** (CRUD operations)
- **Cart & Orders** (Add to cart, checkout, track orders)
- **Swagger API Documentation**
- **Secure with Middleware (Auth & Role-based Access Control)**

---

## 🛠 Tech Stack
- **Backend:** Go (Gin framework)
- **Database:** PostgreSQL (GORM ORM)
- **Authentication:** JWT (JSON Web Tokens)
- **API Documentation:** Swagger

---

## 📌 Installation & Setup
### 1️⃣ Clone the Repository
```sh
git clone https://github.com/telman/ecom-go.git
cd ecom
```

### 2️⃣ Install Dependencies
```sh
go mod tidy
```

### 3️⃣ Setup Environment Variables
Create a `.env` file and add:
```env
DB_HOST=localhost
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=ecom_db
DB_PORT=5432
JWT_SECRET=your_secret_key
```

### 4️⃣ Run Migrations
```sh
go run main.go migrate
```

### 5️⃣ Start the Server
```sh
go run main.go
```

Server will run on `http://localhost:8080`

---

## 📖 API Documentation (Swagger)
Swagger UI is available at:
```
http://localhost:8080/swagger/index.html
```

To regenerate Swagger docs, run:
```sh
swag init
```

---

## 🔥 API Endpoints
### **Auth Routes**
| Method | Endpoint         | Description |
|--------|----------------|-------------|
| POST   | /auth/register  | Register a new user |
| POST   | /auth/login     | Login user and get JWT |
| GET    | /user/profile   | Get user profile |

### **Product Routes**
| Method | Endpoint       | Description |
|--------|---------------|-------------|
| POST   | /products/add  | Add a new product (Admin) |
| GET    | /products      | Get all products |
| GET    | /products/:id  | Get product details |
| PUT    | /products/:id  | Update product (Admin) |
| DELETE | /products/:id  | Delete product (Admin) |

### **Order & Cart Routes**
| Method | Endpoint        | Description |
|--------|----------------|-------------|
| POST   | /cart/add       | Add product to cart |
| GET    | /cart           | View cart items |
| POST   | /orders/place   | Place an order |
| GET    | /orders         | Get user orders |
| PATCH  | /orders/:id/status | Update order status (Admin) |

---

## 📂 Project Structure
```
/ecom-backend
│── controllers/    # API handlers
│── middleware/     # JWT authentication middleware
│── models/         # Database models (GORM)
│── routes/         # API route definitions
│── db/             # Database connection
│── main.go         # Entry point
│── go.mod          # Go module file
│── swagger.json    # API documentation
```

---

## 🔥 Contributing
Feel free to fork and contribute! Open an issue or submit a PR.

---

## 📜 License
This project is licensed under the MIT License. 

---

Happy Coding! 🚀