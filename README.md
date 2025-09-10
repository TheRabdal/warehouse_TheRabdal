# Simple Warehouse Management System

This is a full-stack web application for basic inventory tracking, built as a solution for a technical test case. It features a Vue.js front-end, a Go back-end, and a MySQL database.

---

## Tech Stack

*   **Front-end:** Vue.js (v3)
*   **Back-end:** Go (Golang)
*   **Database:** MySQL
*   **API Communication:** RESTful API (JSON)
*   **Go Dependencies:**
    *   `gorilla/mux` for routing.
    *   `go-sql-driver/mysql` for database connection.
    *   `golang.org/x/crypto/bcrypt` for password hashing.
*   **Front-end Dependencies:**
    *   `axios` for API requests.

---

## Features

*   **Secure User Authentication:** Login system with hashed passwords stored in the database.
*   **Product Management (CRUD):**
    *   **Create:** Add new products to the inventory.
    *   **Read:** View a list of all products from the database.
    *   **Update:** Modify the status of existing products.
    *   **Delete:** Remove products from the inventory.
*   **RESTful API:** A well-defined API to handle all operations.

---

## Prerequisites

Before you begin, ensure you have the following installed:

*   [Node.js](https://nodejs.org/) (which includes npm)
*   [Go](https://golang.org/)
*   [MySQL](https://www.mysql.com/) or MariaDB

---

## Setup and Installation

Follow these steps to get the application running on your local machine.

**1. Database Setup**

*   Make sure your MySQL server is running.
*   Using a MySQL client (like MySQL Workbench, DBeaver, etc.), run the script located at `backend/init.sql`.
    ```bash
    # Example using the command line from the 'backend' directory:
    mysql -u YOUR_USERNAME -p < init.sql
    ```
*   This will create the `warehouse_db` database and the required `products` and `users` tables.

**2. Back-end Setup**

*   Navigate to the `backend` directory:
    ```bash
    cd backend
    ```
*   **Important:** Open the `main.go` file and modify the database connection string with your actual MySQL username and password:
    ```go
    // Replace "root:admin123" with your credentials
    db, err = sql.Open("mysql", "YOUR_USERNAME:YOUR_PASSWORD@tcp(127.0.0.1:3306)/warehouse_db")
    ```
*   Run the back-end server:
    ```bash
    go run main.go
    ```
*   The server will start on `http://localhost:8080`.

**3. Front-end Setup**

*   In a **new, separate terminal**, navigate to the project's root directory (`warehouse`).
*   Install the necessary npm packages:
    ```bash
    npm install
    ```
*   Run the front-end development server:
    ```bash
    npm run serve
    ```
*   The application will be available at the URL provided (usually `http://localhost:8081`).

**4. Login**

*   Open the application in your browser.
*   Use the following credentials to log in:
    *   **Username:** `admin`
    *   **Password:** `admin123`

---

## API Documentation

| Method | Endpoint                  | Description                                      |
|--------|---------------------------|--------------------------------------------------|
| `POST` | `/api/login`              | Authenticates a user.                            |
| `GET`  | `/api/products`           | Retrieves a list of all products.                |
| `POST` | `/api/products`           | Adds a new product to the inventory.             |
| `PUT`  | `/api/products/{sku}`     | Updates the details of an existing product.      |
| `DELETE`| `/api/products/{sku}`     | Deletes a product from the inventory.            |"# Warehouse_project" 
"# warehouse_TheRabdal" 
