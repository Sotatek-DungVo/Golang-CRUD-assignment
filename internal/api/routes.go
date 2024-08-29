package api

import (
	"social-sys/internal/api/handlers"
	"social-sys/internal/repository"
	"social-sys/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handlers.NewPostHandler(postService)

	postRoutes := router.Group("/posts")
	{
		postRoutes.POST("", postHandler.CreatePost)
		postRoutes.GET("", postHandler.ListPosts)
		postRoutes.GET("/:id", postHandler.GetPost)
		postRoutes.PUT("/:id", postHandler.UpdatePost)
		postRoutes.DELETE("/:id", postHandler.DeletePost)
	}

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo);
	authHandler := handlers.NewAuthHandler(authService)
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}
}
