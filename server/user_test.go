package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"users/domain"
	userMiddleware "users/middlewares/user"
	"users/mocks"
	userService "users/services/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockUserRepo := new(mocks.UserMockRepository)

	user := domain.User{
		Username:  "poloyacero10",
		Email:     "akosipoloyacero10@gmail.com",
		Birthdate: "1987-12-18",
	}
	jsonValue, _ := json.Marshal(user)

	mockUserRepo.On("Save").Return(&user, nil)
	userTestService := userService.NewService(mockUserRepo)
	userMiddleware := userMiddleware.NewMiddleware(mockUserRepo)
	handler := userHandler{userTestService, userMiddleware}

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("user", user)
	})
	r.POST("/users", handler.Create)

	reqFound, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdate(t *testing.T) {
	mockUserRepo := new(mocks.UserMockRepository)

	userID := uuid.NewString()

	user := domain.User{
		ID:        userID,
		Birthdate: "1987-12-18",
	}
	updateUserReq := domain.UpdateUserRequest{
		Birthdate: "1987-12-18",
	}
	jsonValue, _ := json.Marshal(updateUserReq)

	mockUserRepo.On("Update").Return(&user, nil)
	userTestService := userService.NewService(mockUserRepo)
	userMiddleware := userMiddleware.NewMiddleware(nil)
	handler := userHandler{userTestService, userMiddleware}

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("user", updateUserReq)
	})

	r.PUT("/users/:id", handler.Update)
	reqFound, _ := http.NewRequest("PUT", "/users/"+userID, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDelete(t *testing.T) {
	mockUserRepo := new(mocks.UserMockRepository)

	user := domain.User{
		ID:        uuid.NewString(),
		Username:  "poloyacero",
		Email:     "akosipoloyacero@gmail.com",
		Birthdate: "1987-12-18",
	}

	mockUserRepo.On("Delete").Return(nil)
	userTestService := userService.NewService(mockUserRepo)
	userMiddleware := userMiddleware.NewMiddleware(nil)
	handler := userHandler{userTestService, userMiddleware}

	r := gin.Default()
	r.DELETE("/users/:id", handler.Delete)
	reqFound, _ := http.NewRequest("DELETE", "/users/"+user.ID, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRead(t *testing.T) {
	mockUserRepo := new(mocks.UserMockRepository)

	user := domain.User{
		ID:        uuid.NewString(),
		Username:  "poloyacero",
		Email:     "akosipoloyacero@gmail.com",
		Birthdate: "1987-12-18",
	}

	mockUserRepo.On("Find").Return(&user, nil)
	userTestService := userService.NewService(mockUserRepo)
	userMiddleware := userMiddleware.NewMiddleware(nil)
	handler := userHandler{userTestService, userMiddleware}

	r := gin.Default()
	r.GET("/users/:id", handler.Read)
	reqFound, _ := http.NewRequest("GET", "/users/"+user.ID, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

	reqNotFound, _ := http.NewRequest("GET", "/users/100", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestStatus(t *testing.T) {
	userService := userService.NewService(nil)
	userMiddleware := userMiddleware.NewMiddleware(nil)

	handler := userHandler{userService, userMiddleware}

	mockResponse := `{"success":true}`
	r := gin.Default()

	r.GET("/users/status", handler.Status)
	req, _ := http.NewRequest("GET", "/users/status", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
