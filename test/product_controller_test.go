package test

// import (
// 	"context"
// 	"encoding/json"
// 	"golang-point-of-sales-system/app"
// 	"golang-point-of-sales-system/controller"
// 	"golang-point-of-sales-system/middleware"
// 	"golang-point-of-sales-system/modules/products/domain/entity"
// 	"golang-point-of-sales-system/modules/products/domain/repository"
// 	"golang-point-of-sales-system/modules/products/domain/service"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/go-playground/validator/v10"
// 	"github.com/google/uuid"
// 	"github.com/jmoiron/sqlx"
// 	"github.com/stretchr/testify/assert"
// )

// func setupTestDB() *sqlx.DB {
// 	//connect to a PostgreSQL database
// 	// Replace the connection details (user, dbname, password, host) with your own
// 	db, err := sqlx.Connect("postgres", "user=postgres dbname=point-of-sales-golang-test sslmode=disable password=postgres host=localhost")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	db.SetMaxIdleConns(5)
// 	db.SetMaxOpenConns(20)
// 	db.SetConnMaxLifetime(60 * time.Minute)
// 	db.SetConnMaxIdleTime(10 * time.Minute)

// 	// Test the connection to the database
// 	if err := db.Ping(); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		log.Println("Successfully Connected")
// 	}

// 	return db

// }

// func setupRouter(db *sqlx.DB) http.Handler {
// 	validate := validator.New()
// 	productRepository := repository.NewProductRepository(db)
// 	productService := service.NewProductService(productRepository, validate)
// 	productController := controller.NewProductController(productService)

// 	router := app.NewRouter(productController)

// 	return middleware.NewAuthMiddleware(router)
// }

// func truncateProduct(db *sqlx.DB) {
// 	db.Exec("TRUNCATE product")
// }

// // Id          uuid.UUID `validate:"required" json:"id"`
// // Kode_produk string    `validate:"required" json:"kode_produk"`
// // Nama_produk string    `validate:"required" json:"nama_produk"`
// // Merk        string    `validate:"required" json:"merk"`
// // Harga_beli  int       `validate:"required" json:"harga_beli"`
// // Harga_jual  int       `validate:"required" json:"harga_jual"`
// // Stok        int       `validate:"required" json:"stok"`

// func TestCreateProductSuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateProduct(db)
// 	router := setupRouter(db)

// 	routerBody := strings.NewReader(`{
// 		"kode_produk": "P-001",
// 		"nama_produk": "Product 001",
// 		"merk": "Merk 001",
// 		"harga_beli": 1000,
// 		"harga_jual": 2000,
// 		"stok": 10
// 	}`)

// 	tRequest := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/product/add", routerBody)
// 	tRequest.Header.Add("Content-Type", "application/json")
// 	tRequest.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, tRequest)

// 	response := recorder.Result()
// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])
// 	assert.Equal(t, "Product 001", responseBody["data"].(map[string]interface{})["nama_produk"])
// 	assert.Equal(t, "P-001", responseBody["data"].(map[string]interface{})["kode_produk"])
// 	assert.Equal(t, "Merk 001", responseBody["data"].(map[string]interface{})["merk"])
// 	assert.Equal(t, 1000, int(responseBody["data"].(map[string]interface{})["harga_beli"].(float64)))
// 	assert.Equal(t, 2000, int(responseBody["data"].(map[string]interface{})["harga_jual"].(float64)))
// }
// func TestCreateProductFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateProduct(db)
// 	router := setupRouter(db)

// 	routerBody := strings.NewReader(`{
// 		"kode_produk": "P-001",
// 		"merk": "Merk 001",
// 		"harga_beli": 1000,
// 		"harga_jual": 2000,
// 		"stok": 10
// 	}`)

// 	tRequest := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/product/add", routerBody)
// 	tRequest.Header.Add("Content-Type", "application/json")
// 	tRequest.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, tRequest)

// 	response := recorder.Result()
// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)

// 	assert.Equal(t, 400, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "BAD REQUEST", responseBody["status"])
// }

// func TestUpdateProductSuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateProduct(db)
// 	router := setupRouter(db)

// 	tx := db.MustBegin()
// 	productRepository := repository.NewProductRepository(db)
// 	product := productRepository.Save(context.Background(), entity.Product{
// 		Kode_produk: "P-001",
// 		Nama_produk: "Product 002",
// 		Merk:        "Merk 001",
// 		Harga_beli:  1000,
// 		Harga_jual:  2000,
// 		Stok:        10,
// 	})
// 	tx.Commit()

// 	routerBody := strings.NewReader(`{
// 		"kode_produk": "P-002",
// 		"nama_produk": "Product 002",
// 		"merk": "Merk 001",
// 		"harga_beli": 1000,
// 		"harga_jual": 2000,
// 		"stok": 10
// 	}`)

// 	tRequest := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/product/update/"+product.Id.String(), routerBody)
// 	tRequest.Header.Add("Content-Type", "application/json")
// 	tRequest.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, tRequest)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])
// 	assert.Equal(t, product.Id, uuid.MustParse(responseBody["data"].(map[string]interface{})["id"].(string)))
// 	assert.Equal(t, "Product 002", responseBody["data"].(map[string]interface{})["nama_produk"])
// 	assert.Equal(t, "P-002", responseBody["data"].(map[string]interface{})["kode_produk"])
// 	assert.Equal(t, "Merk 001", responseBody["data"].(map[string]interface{})["merk"])
// 	assert.Equal(t, 1000, int(responseBody["data"].(map[string]interface{})["harga_beli"].(float64)))
// 	assert.Equal(t, 2000, int(responseBody["data"].(map[string]interface{})["harga_jual"].(float64)))
// }

// func TestUpdateProductFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateProduct(db)
// 	router := setupRouter(db)

// 	tx := db.MustBegin()
// 	productRepository := repository.NewProductRepository(db)
// 	product := productRepository.Save(context.Background(), entity.Product{
// 		Kode_produk: "P-001",
// 		Nama_produk: "Product 002",
// 		Merk:        "Merk 001",
// 		Harga_beli:  1000,
// 		Harga_jual:  2000,
// 		Stok:        10,
// 	})
// 	tx.Commit()

// 	routerBody := strings.NewReader(`{
// 		"kode_produk": "P-002",
// 		"merk": "Merk 001",
// 		"harga_beli": 1000,
// 		"harga_jual": 2000,
// 		"stok": 10
// 	}`)

// 	tRequest := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/product/update/"+product.Id.String(), routerBody)
// 	tRequest.Header.Add("Content-Type", "application/json")
// 	tRequest.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, tRequest)

// 	response := recorder.Result()
// 	assert.Equal(t, 400, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 400, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "BAD REQUEST", responseBody["status"])
// }

// func TestGetProductSuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateProduct(db)
// 	router := setupRouter(db)

// 	tx := db.MustBegin()
// 	productRepository := repository.NewProductRepository(db)
// 	product := productRepository.Save(context.Background(), entity.Product{
// 		Kode_produk: "P-001",
// 		Nama_produk: "Product 002",
// 		Merk:        "Merk 001",
// 		Harga_beli:  1000,
// 		Harga_jual:  2000,
// 		Stok:        10,
// 	})
// 	tx.Commit()

// 	tRequest := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/product/show/"+product.Id.String(), nil)
// 	tRequest.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, tRequest)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])
// 	assert.Equal(t, product.Id, uuid.MustParse(responseBody["data"].(map[string]interface{})["id"].(string)))
// 	assert.Equal(t, product.Nama_produk, responseBody["data"].(map[string]interface{})["nama_produk"])
// }

// func TestGetProductFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateProduct(db)
// 	router := setupRouter(db)

// 	tRequest := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/product/show/e09a5e7b-95ed-4889-8170-fa52ab33a9ea", nil)
// 	tRequest.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, tRequest)

// 	response := recorder.Result()
// 	assert.Equal(t, 404, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 404, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "NOT FOUND", responseBody["status"])
// }

// func TestDeleteProductSuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateProduct(db)
// 	router := setupRouter(db)

// 	tx := db.MustBegin()
// 	productRepository := repository.NewProductRepository(db)
// 	product := productRepository.Save(context.Background(), entity.Product{
// 		Kode_produk: "P-001",
// 		Nama_produk: "Product 002",
// 		Merk:        "Merk 001",
// 		Harga_beli:  1000,
// 		Harga_jual:  2000,
// 		Stok:        10,
// 	})
// 	tx.Commit()

// 	tRequest := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/product/delete/"+product.Id.String(), nil)
// 	tRequest.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, tRequest)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])
// }

// func TestDeleteProductFailed(t *testing.T) {
// 	db := setupTestDB()
// 	truncateProduct(db)
// 	router := setupRouter(db)

// 	tRequest := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/product/delete/e09a5e7b-95ed-4889-8170-fa52ab33a9ea", nil)
// 	tRequest.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, tRequest)

// 	response := recorder.Result()
// 	assert.Equal(t, 404, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 404, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "NOT FOUND", responseBody["status"])
// }

// func TestListProductSuccess(t *testing.T) {
// 	db := setupTestDB()
// 	truncateProduct(db)
// 	router := setupRouter(db)

// 	tx := db.MustBegin()
// 	productRepository := repository.NewProductRepository(db)
// 	product := productRepository.Save(context.Background(), entity.Product{
// 		Kode_produk: "P-001",
// 		Nama_produk: "Product 002",
// 		Merk:        "Merk 001",
// 		Harga_beli:  1000,
// 		Harga_jual:  2000,
// 		Stok:        10,
// 	})
// 	product2 := productRepository.Save(context.Background(), entity.Product{
// 		Kode_produk: "P-002",
// 		Nama_produk: "Product 003",
// 		Merk:        "Merk 003",
// 		Harga_beli:  2000,
// 		Harga_jual:  3000,
// 		Stok:        80,
// 	})
// 	tx.Commit()

// 	tRequest := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/product/lists", nil)
// 	tRequest.Header.Add("X-API-Key", "RAHASIA")

// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, tRequest)

// 	response := recorder.Result()
// 	assert.Equal(t, 200, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 200, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "OK", responseBody["status"])

// 	var products = responseBody["data"].([]interface{})

// 	productResponse1 := products[0].(map[string]interface{})
// 	productResponse2 := products[1].(map[string]interface{})

// 	assert.Equal(t, product.Id, uuid.MustParse(productResponse1["id"].(string)))
// 	assert.Equal(t, product.Nama_produk, productResponse1["nama_produk"])

// 	assert.Equal(t, product2.Id, uuid.MustParse(productResponse2["id"].(string)))
// 	assert.Equal(t, product2.Nama_produk, productResponse2["nama_produk"])
// }

// func TestUnauthorized(t *testing.T) {
// 	db := setupTestDB()
// 	truncateProduct(db)
// 	router := setupRouter(db)

// 	tRequest := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/product/lists", nil)
// 	// tRequest.Header.Add("X-API-Key", "SALAH")
// 	recorder := httptest.NewRecorder()

// 	router.ServeHTTP(recorder, tRequest)

// 	response := recorder.Result()
// 	assert.Equal(t, 401, response.StatusCode)

// 	body, _ := io.ReadAll(response.Body)
// 	var responseBody map[string]interface{}
// 	json.Unmarshal(body, &responseBody)
// 	assert.Equal(t, 401, int(responseBody["code"].(float64)))
// 	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
// }
