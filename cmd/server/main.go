package main

import (
	"context"
	"log"
	"vm-crud/internal/database"
	"vm-crud/internal/handler"
	"vm-crud/internal/repository"
	"vm-crud/internal/service"

	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("warning: .env file not found")
	}
	fmt.Println("DB_HOST:", os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT:", os.Getenv("DB_PORT"))
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_NAME:", os.Getenv("DB_NAME"))
	//  define  db  connection here
	ctx := context.Background()
	dbpool, err := database.NewPostgresql(ctx)
	if err != nil {
		log.Fatalf("database  connection  failed  %v", err)
	}
	defer dbpool.Close()
	log.Println("PostgreSQL connected successfully")

	// Repository
	vmRepository := repository.NewVMRepository(dbpool)
	// Service
	vmService := service.NewVMService(vmRepository)

	// Handler
	vmHandler := handler.NewVMHandler(vmService)

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	v1 := router.Group("/v1")
	{
		v1.POST("/vms", vmHandler.CreateVM)
		v1.GET("/vms", vmHandler.ListVMs)
		v1.GET("/vms/:id", vmHandler.GetVMByID)
		v1.PUT("/vms/:id", vmHandler.UpdateVM)
		v1.DELETE("/vms/:id", vmHandler.DeleteVM)
	}
	router.Run(":8080")
}
