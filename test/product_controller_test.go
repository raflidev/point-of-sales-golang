package test

import (
	"golang-point-of-sales-system/app"
	"golang-point-of-sales-system/controller"
	"golang-point-of-sales-system/modules/products/domain/repository"
	"golang-point-of-sales-system/modules/products/domain/service"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sqlx.DB {
	//connect to a PostgreSQL database
	// Replace the connection details (user, dbname, password, host) with your own
	db, err := sqlx.Connect("postgres", "user=postgres dbname=point-of-sales-golang-test sslmode=disable password=postgres host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	return db

}

func setupRouter() *httprouter.Router {
	db := setupTestDB()
	validate := validator.New()
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository, validate)
	productController := controller.NewProductController(productService)

	router := app.NewRouter(productController)

	return router
}

// Id          uuid.UUID `validate:"required" json:"id"`
// Kode_produk string    `validate:"required" json:"kode_produk"`
// Nama_produk string    `validate:"required" json:"nama_produk"`
// Merk        string    `validate:"required" json:"merk"`
// Harga_beli  int       `validate:"required" json:"harga_beli"`
// Harga_jual  int       `validate:"required" json:"harga_jual"`
// Stok        int       `validate:"required" json:"stok"`

func TestCreateProductSuccess(t *testing.T) {
	router := setupRouter()

	routerBody := strings.NewReader(`{
		"kode_produk": "P-001",
		"nama_produk": "Product 001",
		"merk": "Merk 001",
		"harga_beli": 1000,
		"harga_jual": 2000,
		"stok": 10
	}`)

	tRequest := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/product/add", routerBody)
	tRequest.Header.Add("Content-Type", "application/json")
	tRequest.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, tRequest)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}
func TestCreateProductFailed(t *testing.T) {

}

func TestUpdateProductSuccess(t *testing.T) {

}

func TestUpdateProductFailed(t *testing.T) {

}

func TestGetProductSuccess(t *testing.T) {

}

func TestGetProductFailed(t *testing.T) {

}

func TestDeleteProductSuccess(t *testing.T) {

}

func TestDeleteProductFailed(t *testing.T) {

}

func TestListProductSuccess(t *testing.T) {

}

func TestUnauthorized(t *testing.T) {

}
