package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"signoz-test/controllers"
	"signoz-test/dto"
	"signoz-test/metrics"
	"signoz-test/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestController_AddToCart_Success(t *testing.T) {
	metrics.InitMeters()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	controller := controllers.NewController(mockService)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/cart/add", controller.AddToCart)

	cartName := "cart1"
	itemName := "item1"
	cartItem := dto.AddToCart{CartName: &cartName, ItemName: &itemName}
	mockService.EXPECT().AddItemToCart(gomock.Any(), cartItem).Return(nil)

	body, _ := json.Marshal(cartItem)
	req, _ := http.NewRequest(http.MethodPost, "/cart/add", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	var resp dto.Response
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 201, resp.Code)
	assert.Equal(t, "operation_successful", resp.Message)
}

func TestController_AddToCart_BindJSONError(t *testing.T) {
	metrics.InitMeters()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	controller := controllers.NewController(mockService)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/cart/add", controller.AddToCart)

	// Invalid JSON
	body := []byte(`{"cart_name": 123}`)
	req, _ := http.NewRequest(http.MethodPost, "/cart/add", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	var resp dto.Response
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 400, resp.Code)
	assert.NotEmpty(t, resp.Message)
}

func TestController_AddToCart_ValidateError(t *testing.T) {
	metrics.InitMeters()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	controller := controllers.NewController(mockService)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/cart/add", controller.AddToCart)

	// Missing item_name
	cartName := "cart1"
	cartItem := dto.AddToCart{CartName: &cartName, ItemName: nil}
	body, _ := json.Marshal(cartItem)
	req, _ := http.NewRequest(http.MethodPost, "/cart/add", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	var resp dto.Response
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 400, resp.Code)
	assert.NotEmpty(t, resp.Message)
}

func TestController_AddToCart_ServiceError(t *testing.T) {
	metrics.InitMeters()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	controller := controllers.NewController(mockService)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/cart/add", controller.AddToCart)

	cartName := "cart1"
	itemName := "item1"
	cartItem := dto.AddToCart{CartName: &cartName, ItemName: &itemName}
	mockService.EXPECT().AddItemToCart(gomock.Any(), cartItem).Return(errors.New("service error"))

	body, _ := json.Marshal(cartItem)
	req, _ := http.NewRequest(http.MethodPost, "/cart/add", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	var resp dto.Response
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(t, "service error", resp.Message)
}

func TestController_GetItemsInCart_Success(t *testing.T) {
	metrics.InitMeters()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	controller := controllers.NewController(mockService)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/cart/:cartName", controller.GetItemsInCart)

	cartName := "cart1"
	items := &dto.ItemsInCart{Items: []string{"item1", "item2"}}
	itemsCopy := map[string]interface{}{
		"items": []interface{}{
			"item1", "item2",
		},
	}
	mockService.EXPECT().GetItemsInCart(gomock.Any(), cartName).Return(items, nil)

	req, _ := http.NewRequest(http.MethodGet, "/cart/"+cartName, nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	var resp dto.Response
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, itemsCopy, resp.Message)
}

func TestController_GetItemsInCart_ServiceError(t *testing.T) {
	metrics.InitMeters()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockService(ctrl)
	controller := controllers.NewController(mockService)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/cart/:cartName", controller.GetItemsInCart)

	cartName := "cart1"
	mockService.EXPECT().GetItemsInCart(gomock.Any(), cartName).Return(nil, errors.New("service error"))

	req, _ := http.NewRequest(http.MethodGet, "/cart/"+cartName, nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	var resp dto.Response
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(t, "service error", resp.Message)
}
