package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"inherited/internal/middleware"
	"inherited/internal/models"
	"inherited/internal/services"
	"log"
	"net/http"
)

//func CreateAdminUser(ctx *gin.Context) {
//	//Parameter parsing
//	adminUser := models.SysUser{}
//	err := ctx.BindJSON(&adminUser)
//	if err != nil {
//		fmt.Printf("[api.CreateAdminUser], Parameter parsing error")
//	}
//
//	userName := adminUser.UserName
//	password := adminUser.Password
//	if userName == "" || password == "" {
//		log.Printf("userName or password is null")
//
//		return
//	}
//
//	sysUser := new(services.SysUser)
//	err = sysUser.CreateSysUser(adminUser.UserName, adminUser.Password)
//	if err != nil {
//		fmt.Printf("[api.CreateAdminUser], err: %v", err)
//
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"date": "success",
//	})
//
//	return
//}
func GetAdminUser(ctx *gin.Context) {
	//Parameter parsing
	adminUser := models.SysUser{}
	err := ctx.BindJSON(&adminUser)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], Parameter parsing error")
		return
	}
	num := adminUser.Num
	password := adminUser.Password
	if num == 0 || password == "" {
		log.Printf("userName or password is null")

		return
	}

	sysUser := new(services.SysUser)
	userInfo, err := sysUser.GetSysUser(num, password)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], err: %v", err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"date": userInfo,
	})

	return
}

// 得到Money
func CreateSubInfo(ctx *gin.Context) {
	//Parameter parsing
	info := models.Print{}
	err := ctx.BindJSON(&info)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], Parameter parsing error")
		return
	}

	subInfo := new(services.Print)
	err = subInfo.CreateSubInfo(info.Num, info.Money, info.UserName, info.Scheme)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], err: %v", err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})

	return
}

// 校验code
func VerifyCode(ctx *gin.Context) {
	//Parameter parsing
	info := models.Code{}
	err := ctx.BindJSON(&info)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], Parameter parsing error")
		return
	}

	codeInfo := new(services.Code)
	code,err := codeInfo.VerifyCode(info.Code)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], err: %v", err)

		return
	}
	log.Println("code",code.Code, "er", info.Code )
    if code.Code != info.Code{
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "提取码不正确",
		})
		return
	}
	salaryInfo := new(services.Print)
	salary, err := salaryInfo.GetSalaryInfo()
	if err != nil {
		log.Println("[VerifyCode] ", err)
		return
	}
	middleware.ToExcel(salary, ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})

	return
}
