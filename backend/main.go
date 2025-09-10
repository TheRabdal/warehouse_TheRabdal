package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"warehouse/backend/handlers"
	"warehouse/backend/services"
)

var db *sql.DB

// CORS Middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	var err error
	db, err = sql.Open("mysql", "admin:admin123@tcp(127.0.0.1:3306)/warehouse")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	log.Println("Successfully connected to the database")

	// Initialize Services
	authService := services.NewAuthService(db)
	productService := services.NewProductService(db)

	// Initialize Handlers
	authHandler := handlers.NewAuthHandler(authService)
	productHandler := handlers.NewProductHandler(productService)

	// Init router
	r := mux.NewRouter()

	// API Routes
	r.HandleFunc("/api/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/api/register", authHandler.Register).Methods("POST")

	// Product CRUD
	r.HandleFunc("/api/products", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/api/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/api/products/{sku}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/products/{sku}", productHandler.DeleteProduct).Methods("DELETE")

	// Product CSV Export
	r.HandleFunc("/api/products/export/csv", productHandler.ExportProductsCSV).Methods("GET")

	http.Handle("/", corsMiddleware(r))

	log.Println("Server starting on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
