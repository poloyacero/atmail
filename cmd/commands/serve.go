package commands

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	userMiddleware "users/middlewares/user"
	"users/server"
	userService "users/services/user"
	"users/storage"
	"users/storage/mysql"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run http server",
	RunE: func(cmd *cobra.Command, args []string) error {
		database := storage.NewDatabase()
		db, err := database.New()
		if err != nil {
			return err
		}

		// Set up repositories
		usersRepository := mysql.NewSourceRepository(db)

		// Set up services
		userService := userService.NewService(usersRepository)
		userMiddleware := userMiddleware.NewMiddleware(usersRepository)

		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			_ = v.RegisterValidation("validDate", userMiddleware.ValidDateFunc)
			_ = v.RegisterValidation("unique", userMiddleware.Unique)
		}

		router := gin.Default()

		{
			group := router.Group("/users", gin.BasicAuth(gin.Accounts{
				"admin": "opensesame",
			}))
			server.NewUserHandler(group, userService, userMiddleware)
		}

		srv := &http.Server{
			Addr:    ":8080",
			Handler: router,
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		quit := make(chan os.Signal)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server forced to shutdown:", err)
		}

		log.Println("Server exiting")
		return nil
	},
}
