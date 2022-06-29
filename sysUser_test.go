package main

import (
	"fmt"
	"inherited/internal"
	"inherited/internal/services"
	"log"
	"testing"
)

func TestCreateSysUser(t *testing.T) {

	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	SysUser := new(services.SysUser)
	user, _ := SysUser.GetSysUser(199020440236, "2355")
	fmt.Println("er", user)

}

func TestUser(t *testing.T) {
	fmt.Println("er")
}

func TestCreateSubInfo(t *testing.T) {

	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	SysUser := new(services.Print)
<<<<<<< HEAD
	_ = SysUser.CreateSubInfo(199020440236, 2000000, "dong", "A")

}

func TestGetTableInfo(t *testing.T) {

	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	SysUser := new(services.Print)
	info, _ := SysUser.GetSalaryInfo()

	fmt.Println("answer", info)
=======
	_ = SysUser.CreateSubInfo(199020440236, 2000000)
>>>>>>> 81203b5 (选择A，B)

}
