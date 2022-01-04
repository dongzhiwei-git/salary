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
	_ = SysUser.CreateSubInfo(199020440236, 2000000)

}
