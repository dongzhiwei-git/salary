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
<<<<<<< HEAD
		admin.POST("/login")
		admin.POST("/get-table", api.VerifyCode)
=======
		admin.POST("/login", )
>>>>>>> 81203b5 (选择A，B)
	}

	// setup listen
	err := r.Run(":26667")
	if err != nil {
		fmt.Printf("run failed: %v\n", err)

		return


	}

}
