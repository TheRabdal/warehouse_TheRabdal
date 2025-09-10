package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"encoding/csv"
	"warehouse/backend/models"
	"warehouse/backend/services"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(ps *services.ProductService) *ProductHandler {
	return &ProductHandler{productService: ps}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	statusFilter := r.URL.Query().Get("status")

	products, err := h.productService.GetProducts(statusFilter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p models.Product
	_ = json.NewDecoder(r.Body).Decode(&p)

	if err := h.productService.CreateProduct(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	sku := params["sku"]
	var p models.Product
	_ = json.NewDecoder(r.Body).Decode(&p)

	if err := h.productService.UpdateProduct(sku, &p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(p)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	sku := params["sku"]

	if err := h.productService.DeleteProduct(sku); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) ExportProductsCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=products.csv")

	writer := csv.NewWriter(w)
	defer writer.Flush()

	products, err := h.productService.GetProducts("") // Get all products
	if err != nil {
		http.Error(w, "Failed to fetch products for CSV export: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for _, product := range products {
		row := []string{
			product.SKU,
			product.Name,
			strconv.Itoa(product.Quantity),
			product.Location,
			product.Status,
		}
		if err := writer.Write(row); err != nil {
			http.Error(w, "Failed to write CSV row: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}