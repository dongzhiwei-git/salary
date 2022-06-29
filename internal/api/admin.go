package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
<<<<<<< HEAD
<<<<<<< HEAD
	"inherited/internal/middleware"
=======
>>>>>>> 81203b5 (选择A，B)
=======
	"inherited/internal/middleware"
>>>>>>> d73ad65 (基本完成)
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
	err := ctx.ShouldBindJSON(&adminUser)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], Parameter parsing error")
<<<<<<< HEAD
<<<<<<< HEAD
		ctx.JSON(http.StatusOK, gin.H{
			"date": "格式不对",
		})
=======
>>>>>>> 81203b5 (选择A，B)
=======
		ctx.JSON(http.StatusOK, gin.H{
			"date": "格式不对",
		})
>>>>>>> be39e95 (修改一些gorm)
		return
	}
	num := adminUser.Num
	password := adminUser.Password
	if num == 0 || password == "" {
		log.Printf("userName or password is null")

		return
	}

	sysUser := new(services.SysUser)
<<<<<<< HEAD
<<<<<<< HEAD
	adminInfo, err := sysUser.GetSysUser(num, password)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], err: %v", err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "账号或密码不对",
		})
=======
	userInfo, err := sysUser.GetSysUser(num, password)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], err: %v", err)

>>>>>>> 81203b5 (选择A，B)
=======
	adminInfo, err := sysUser.GetSysUser(num, password)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], err: %v", err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "账号或密码不对",
		})
>>>>>>> be39e95 (修改一些gorm)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
<<<<<<< HEAD
<<<<<<< HEAD
		"msg": "登录成功",
		"data": adminInfo,
=======
		"date": userInfo,
>>>>>>> 81203b5 (选择A，B)
=======
		"msg": "登录成功",
		"data": adminInfo,
>>>>>>> be39e95 (修改一些gorm)
	})

	return
}

// 得到Money
func CreateSubInfo(ctx *gin.Context) {
	//Parameter parsing
	info := models.Print{}
	err := ctx.BindJSON(&info)
	if err != nil {
<<<<<<< HEAD
<<<<<<< HEAD
		fmt.Println("[api.CreateAdminUser], Parameter parsing error", err)
=======
		fmt.Printf("[api.CreateAdminUser], Parameter parsing error")
>>>>>>> 81203b5 (选择A，B)
=======
		fmt.Println("[api.CreateAdminUser], Parameter parsing error", err)
>>>>>>> be39e95 (修改一些gorm)
		return
	}

	subInfo := new(services.Print)
<<<<<<< HEAD
<<<<<<< HEAD
	err = subInfo.CreateSubInfo(info.Num, info.Money, info.UserName, info.Scheme)
=======
	err = subInfo.CreateSubInfo(info.Num, info.Money)
>>>>>>> 81203b5 (选择A，B)
=======
	err = subInfo.CreateSubInfo(info.Num, info.Money, info.UserName, info.Scheme)
>>>>>>> d73ad65 (基本完成)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], err: %v", err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
<<<<<<< HEAD
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
=======
>>>>>>> 81203b5 (选择A，B)
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
