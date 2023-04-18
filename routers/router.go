package routers

import (
	"sesi_12/controllers"
	"sesi_12/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", controllers.CreateProduct)
		productRouter.PUT("/:productId", middlewares.ProductAuthorizationOnDeletUpdate(), controllers.UpdateProduct)
		productRouter.GET("/:productId", middlewares.ProductAuthorization(), controllers.GetProductById)
		productRouter.DELETE("/:productId", middlewares.ProductAuthorizationOnDeletUpdate(), controllers.DeleteProductById)
	}
	return r
}
