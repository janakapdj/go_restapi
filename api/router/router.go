package router

import (
	"mytest.net/restapi/api/controller"
	"mytest.net/restapi/api/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Controller Instance
var ctl = controller.GetController()

// App struct coontains
type App struct {
	clients *model.Clients
}

//SetupRouter func
func SetupRouter(c *model.Clients) *gin.Engine {
	app := new(App)
	app.clients = c
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE", "PATCH"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Access-Control-Allow-Origin", "Cache-Control", "Cookie", "Accept", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Token", "Refresh-Token", "Expire"},
		AllowCredentials: true,
		AllowWildcard:    true,
	}))
	v2 := router.Group("/api/v1")
	{
		products := v2.Group("/products")
		{
			products.GET("",  app.handleRequest(ctl.GetProducts))
			products.POST("/create",app.handleRequest(ctl.CreateProduct))
			products.POST("/byId", app.handleRequest(ctl.GetProductByID))
			products.PUT("/update/:id", app.handleRequest(ctl.UpdateProduct))
			products.DELETE(":id", app.handleRequest(ctl.DeleteProduct))
		}
		sales := v2.Group("/sales")
		{
			sales.GET("",  app.handleRequest(ctl.GetSales))
			sales.POST("/create",app.handleRequest(ctl.CreateSale))
			sales.POST("/byId", app.handleRequest(ctl.GetSaleByID))
			sales.PUT("/update/:id", app.handleRequest(ctl.UpdateSale))
			sales.POST("/salesOrderDetails",  app.handleRequest(ctl.SalesOrderDetails))
		}
	}

	return router
}

// RequestHandlerFunction is a custome type that help us to pass db arg to all endpoints
type RequestHandlerFunction func(app *model.Clients, r *gin.Context)

// handleRequest is a middleware we create for pass in db connection to endpoints.
func (app *App) handleRequest(handler RequestHandlerFunction) gin.HandlerFunc {
	return func(r *gin.Context) {
		handler(app.clients, r)
	}
}
