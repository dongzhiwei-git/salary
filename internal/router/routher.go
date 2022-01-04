package router

import (
	"fmt"
	"inherited/internal/api"
	"inherited/internal/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	var r *gin.Engine
	r = gin.Default()
	// to solve the cross domain
	r.Use(pkg.Cors())
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})

	r.POST("/login", api.GetAdminUser)
	r.POST("/sub", api.CreateSubInfo)
	var admin = r.Group("/admin")
	{
		admin.POST("/login", )
	}

	// setup listen
	err := r.Run(":8000")
	if err != nil {
		fmt.Printf("run failed: %v\n", err)
		return

	}

}
