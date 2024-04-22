package server

import (
	"net/http"
	"users/domain"
	userMiddleware "users/middlewares/user"
	"users/services/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userHandler struct {
	userService    user.Service
	userMiddleware userMiddleware.Middleware
}

func (h *userHandler) Create(ctx *gin.Context) {
	request := domain.User{}

	user := ctx.MustGet("user").(domain.User)

	request.ID = uuid.New().String()
	request.Email = user.Email
	request.Username = user.Username
	request.Birthdate = user.Birthdate

	result, err := h.userService.CreateUsers(ctx, request)
	if err != nil {
		CreateError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (h *userHandler) Read(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		CreateError(ctx, http.StatusBadRequest, err)
		return
	}

	result, err := h.userService.FindUser(ctx, userID)
	if err != nil {
		CreateError(ctx, http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (h *userHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		CreateError(ctx, http.StatusBadRequest, err)
		return
	}

	err = h.userService.DeleteUser(ctx, userID)
	if err != nil {
		CreateError(ctx, http.StatusNotFound, err)
		return
	}

	CreateSuccess(ctx, "Success")
}

func (h *userHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		CreateError(ctx, http.StatusBadRequest, err)
		return
	}

	user := ctx.MustGet("user").(domain.UpdateUserRequest)

	request := domain.User{}
	request.ID = uuid.New().String()
	request.Birthdate = user.Birthdate

	result, err := h.userService.UpdateUsers(ctx, request, userID)
	if err != nil {
		CreateError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (h *userHandler) Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, domain.UserSuccessResponse{Success: true})
}

func NewUserHandler(h *gin.RouterGroup, userService user.Service, userMiddleware userMiddleware.Middleware) {
	handler := userHandler{userService, userMiddleware}
	h.GET("/status", handler.Status)
	h.POST("", userMiddleware.CreateMiddleware, handler.Create)
	h.GET("/:id", handler.Read)
	h.PUT("/:id", userMiddleware.UpdateMiddleware(), handler.Update)
	h.DELETE("/:id", handler.Delete)
}
