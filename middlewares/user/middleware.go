package user

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
	"users/domain"
	"users/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MyReadCloser struct {
	rc io.ReadCloser
	w  io.Writer
}

func (rc *MyReadCloser) Read(p []byte) (n int, err error) {
	n, err = rc.rc.Read(p)
	log.Println("run here", n, err)
	if n > 0 {
		if n, err := rc.w.Write(p[:n]); err != nil {
			return n, err
		}
	}
	return n, err
}

func (rc *MyReadCloser) Close() error {
	return rc.rc.Close()
}

type Middleware struct {
	userRepository domain.UserRepository
}

func (m *Middleware) ValidDateFunc(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	if _, err := time.Parse("2006-01-02", input); err != nil {
		return false
	}

	return true
}

func (m *Middleware) Unique(fl validator.FieldLevel) bool {
	input := fl.Field().String()

	if ok := m.userRepository.IsExist(context.TODO(), strings.ToLower(fl.FieldName()), input); !ok {
		return false
	}

	return true
}

func (m *Middleware) UpdateMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := domain.UpdateUserRequest{}

		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				out := make([]utils.ErrorMsg, len(ve))
				for i, fe := range ve {
					out[i] = utils.ErrorMsg{strings.ToLower(fe.Field()), utils.GetErrorMsg(fe)}
				}
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
			}
			return
		}
		// Call the next middleware function
		ctx.Set("user", request)
		ctx.Next()
	}
}

func (m *Middleware) CreateMiddleware(ctx *gin.Context) {
	request := domain.User{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]utils.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = utils.ErrorMsg{strings.ToLower(fe.Field()), utils.GetErrorMsg(fe)}
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	ctx.Set("user", request)
	ctx.Next()
}

func NewMiddleware(userRepository domain.UserRepository) Middleware {
	return Middleware{userRepository}
}
