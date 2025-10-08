package main

import (
    "log"
    "task-manager/internal/handlers"
    "task-manager/internal/middleware"
    "task-manager/internal/storage"

    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize storage
    taskStorage := storage.NewMemoryStorage()

    // // Initialize handlers
    // taskHandler := handlers.NewTaskHandler(taskStorage)

    // // Create Gin router
    // router := gin.Default()

    // // Add middleware
    // router.Use(middleware.Logger())
    // router.Use(middleware.CORSMiddleware())

    // // Health check endpoint
    // router.GET("/health", func(c *gin.Context) {
    //     c.JSON(200, gin.H{
    //         "status":  "healthy",
    //         "service": "task-manager-api",
    //     })
    // })

    // // API routes
    // api := router.Group("/api/v1")
    // {
    //     tasks := api.Group("/tasks")
    //     {
    //         tasks.POST("", taskHandler.CreateTask)
    //         tasks.GET("", taskHandler.GetAllTasks)
    //         tasks.GET("/status", taskHandler.GetTasksByStatus)
    //         tasks.GET("/:id", taskHandler.GetTask)
    //         tasks.PUT("/:id", taskHandler.UpdateTask)
    //         tasks.DELETE("/:id", taskHandler.DeleteTask)
    //     }
    // }

    // Start server
    log.Println("🚀 Task Manager API server starting on :8080")
    log.Println("📚 Available endpoints:")
    log.Println("   GET    /health")
    log.Println("   POST   /api/v1/tasks")
    log.Println("   GET    /api/v1/tasks")
    log.Println("   GET    /api/v1/tasks/status?status=:status")
    log.Println("   GET    /api/v1/tasks/:id")
    log.Println("   PUT    /api/v1/tasks/:id")
    log.Println("   DELETE /api/v1/tasks/:id")

    // if err := router.Run(":8080"); err != nil {
    //     log.Fatal("Failed to start server:", err)
    // }
}