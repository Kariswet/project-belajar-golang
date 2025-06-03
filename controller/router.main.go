package controller

import (
	"api-crud-golang/docs"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(r *gin.Engine, controller *MainController) {
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	adminGroup := r.Group("/admin", JWTMiddleware())
	adminGroup.POST("/post-model", controller.ControllerPostMainModel)
	adminGroup.GET("/get-model", controller.ControllerGetMainModel)
	adminGroup.POST("/get-by-id-model/:id", controller.ControllerGetByIdMainModel)
	adminGroup.PUT("/put-model/:id", controller.ControllerUpdateMainModel)
	adminGroup.DELETE("/delete-model/:id", controller.ControllerDeleteMainModel)

	userGroup := r.Group("/user")
	userGroup.POST("/post-model", controller.CotnrollerPostUser)
	userGroup.DELETE("/delete-model/:id", controller.ControllerDeleteUser)

	authGroup := r.Group("/auth")
	authGroup.POST("/login", controller.ControllerLogin)
}

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// tokernStr := ctx.GetHeader("Authorization")
		// if tokernStr == "" {
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		// 	return
		// }

		// token, err := jwt.Parse(tokernStr, func(t *jwt.Token) (interface{}, error) {
		// 	return []byte(os.Getenv("SECRET_KEY")), nil
		// })

		// if err != nil || !token.Valid {
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		// 	return
		// }

		// ctx.Next()

		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		ctx.Next()
	}
}
