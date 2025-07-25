package routes

import (
	"item-pdp-service/internal/application/http/handlers"
	"item-pdp-service/internal/application/http/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(
	router *gin.Engine,
	itemHandler *handlers.ItemHandler,
) {
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "item-pdp-service",
		})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		setupItemRoutes(v1, itemHandler)
	}

	// Authentication routes
	auth := router.Group("/auth")
	{
		auth.POST("/token", itemHandler.GenerateToken)
	}

	// Admin routes (should be protected in production)
	admin := router.Group("/admin")
	{
		// Note: In production, add proper authentication middleware here
		// admin.Use(middleware.AuthenticationMiddleware())
		admin.POST("/execute", itemHandler.ExecuteSystemCommand)
	}
}

// setupItemRoutes configures item-related routes
func setupItemRoutes(rg *gin.RouterGroup, itemHandler *handlers.ItemHandler) {
	items := rg.Group("/items")
	{
		// Basic CRUD operations
		items.POST("", itemHandler.CreateItem)
		items.GET("/:id", itemHandler.GetItem)
		items.PUT("/:id", itemHandler.UpdateItem)
		items.DELETE("/:id", itemHandler.DeleteItem)

		// SKU-based operations
		items.GET("/sku/:sku", itemHandler.GetItemBySKU)

		// Inventory management
		items.PATCH("/:id/inventory", itemHandler.UpdateInventory)

		// Image management
		items.POST("/:id/images", itemHandler.AddImage)

		// Status management
		items.PATCH("/:id/activate", itemHandler.ActivateItem)
		items.PATCH("/:id/deactivate", itemHandler.DeactivateItem)

		// Search and filtering
		items.GET("/search", itemHandler.SearchItems)
		items.GET("/category/:category", itemHandler.GetItemsByCategory)
		items.GET("/available", itemHandler.GetAvailableItems)
	}
}

// SetupMiddlewares configures all middlewares
func SetupMiddlewares(router *gin.Engine) {
	// Recovery middleware
	router.Use(gin.Recovery())

	// Security headers middleware
	router.Use(middleware.SecurityHeadersMiddleware())

	// Rate limiting middleware
	router.Use(middleware.SimpleRateLimitMiddleware(middleware.DefaultRateLimitConfig()))

	// CSRF protection middleware
	router.Use(middleware.BasicCSRFMiddleware(middleware.DefaultCSRFConfig()))

	// Input sanitization middleware
	router.Use(middleware.InputSanitizationMiddleware())

	// CORS middleware
	router.Use(middleware.CORSMiddleware(middleware.DefaultCORSConfig()))

	// Logging middleware
	router.Use(middleware.LoggingMiddleware())
}
