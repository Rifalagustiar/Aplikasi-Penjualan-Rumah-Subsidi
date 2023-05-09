package route

import (
	"crud/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// Group route dengan JWT middleware
	ejwt := e.Group("/jwt")
	// ejwt.Use(middleware.JWTMiddleware)

	ejwt.GET("/users", controllers.GetUsersController)
	ejwt.GET("/users/:id", controllers.GetUserController)
	ejwt.POST("/users", controllers.CreateUserController)
	ejwt.DELETE("/users/:id", controllers.DeleteUserController)
	ejwt.PUT("/users/:id", controllers.UpdateUserController)

	e.POST("/login", controllers.LoginUserController)

	ejwt.POST("/properties", controllers.CreatePropertyEndpoint)
	ejwt.GET("/properties/:id", controllers.GetPropertyEndpoint)
	ejwt.PUT("/properties/:id", controllers.UpdatePropertyEndpoint)
	ejwt.DELETE("/properties/:id", controllers.DeletePropertyEndpoint)
	ejwt.GET("/properties", controllers.ListPropertiesEndpoint)

	ejwt.GET("/customers", controllers.GetCustomers)
	ejwt.GET("/customers/:id", controllers.GetCustomerByID)
	ejwt.POST("/customers", controllers.CreateCustomer)
	ejwt.PUT("/customers/:id", controllers.UpdateCustomer)
	ejwt.DELETE("/customers/:id", controllers.DeleteCustomer)

	ejwt.GET("/agents", controllers.GetAgentsController)
	ejwt.POST("/agents", controllers.CreateAgentController)
	ejwt.PUT("/agents", controllers.UpdateAgentController)
	ejwt.DELETE("/agents/:id", controllers.DeleteAgentController)

	ejwt.GET("/transactions", controllers.GetTransactions)
	ejwt.GET("/transactions/:id", controllers.GetTransaction)
	ejwt.POST("/transactions", controllers.CreateTransaction)
	ejwt.PUT("/transactions/:id", controllers.UpdateTransaction)
	ejwt.DELETE("/transactions/:id", controllers.DeleteTransaction)

	return e
}
