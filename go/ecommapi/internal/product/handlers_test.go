package product

import (
	"bytes"
	"ecommapi/internal/auth"
	db "ecommapi/internal/helpers/database"
	notif "ecommapi/internal/notification"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.GET("/products", GetProducts)
	r.POST("/products", AddProduct)

	return r
}

func performRequest(router *gin.Engine, method string, path string, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestGetProducts(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name              string
		mockFetchProducts func(_ string, _ int) ([]db.Product, error)
		requestBody       string
		path              string
		expectedStatus    int
	}{
		{
			name:              "success",
			mockFetchProducts: func(_ string, _ int) ([]db.Product, error) { return []db.Product{}, nil },
			path:              "/products?page=1&category_id=test",
			expectedStatus:    http.StatusOK,
		},
		{
			name:           "invalid request",
			path:           "/products?nothing=here",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:              "get products fail",
			mockFetchProducts: func(_ string, _ int) ([]db.Product, error) { return nil, errors.New("error") },
			path:              "/products?page=1&category_id=test",
			expectedStatus:    http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFetchProducts != nil {
				fetchProducts = tt.mockFetchProducts
				defer func() { fetchProducts = FetchProducts }()
			}

			w := performRequest(router, http.MethodGet, tt.path, tt.requestBody)
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestAddProduct(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name                  string
		mockCreateProduct     func(_ ProductDTO) (*db.Product, error)
		mockGetSessionByToken func(_ string) (*db.Session, error)
		mockNotif             func(_, _, _ string) error
		requestBody           string
		expectedStatus        int
	}{
		{
			name:                  "success",
			mockCreateProduct:     func(_ ProductDTO) (*db.Product, error) { return &db.Product{}, nil },
			mockGetSessionByToken: func(_ string) (*db.Session, error) { return &db.Session{}, nil },
			mockNotif:             func(_, _, _ string) error { return nil },
			requestBody:           `{"name":"pname","price":12,"description":"desc","category_id":"cid"}`,
			expectedStatus:        http.StatusCreated,
		},
		{
			name:              "create product fails",
			mockCreateProduct: func(_ ProductDTO) (*db.Product, error) { return nil, errors.New("error") },
			requestBody:       `{"name":"pname","price":12,"description":"desc","category_id":"cid"}`,
			expectedStatus:    http.StatusInternalServerError,
		},
		{
			name:                  "get session fails",
			mockCreateProduct:     func(_ ProductDTO) (*db.Product, error) { return &db.Product{}, nil },
			mockGetSessionByToken: func(_ string) (*db.Session, error) { return nil, errors.New("error") },
			requestBody:           `{"name":"pname","price":12,"description":"desc","category_id":"cid"}`,
			expectedStatus:        http.StatusInternalServerError,
		},
		{
			name:                  "success",
			mockCreateProduct:     func(_ ProductDTO) (*db.Product, error) { return &db.Product{}, nil },
			mockGetSessionByToken: func(_ string) (*db.Session, error) { return &db.Session{}, nil },
			mockNotif:             func(_, _, _ string) error { return errors.New("error") },
			requestBody:           `{"name":"pname","price":12,"description":"desc","category_id":"cid"}`,
			expectedStatus:        http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockCreateProduct != nil {
				createProduct = tt.mockCreateProduct
				defer func() { createProduct = CreateProduct }()
			}
			if tt.mockGetSessionByToken != nil {
				getSessionByToken = tt.mockGetSessionByToken
				defer func() { getSessionByToken = auth.GetSessionByToken }()
			}
			if tt.mockNotif != nil {
				createEventNotif = tt.mockNotif
				defer func() { createEventNotif = notif.CreateEventNotif }()
			}

			w := performRequest(router, http.MethodPost, "/products", tt.requestBody)
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
