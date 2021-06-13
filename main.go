package main

import (
	"fmt"
	check "github.com/asaskevich/govalidator"
	"github.com/DeniesKresna/jobhunop/Configs"
	"github.com/DeniesKresna/jobhunop/Routers"
)

func main() {
	check.SetFieldsRequiredByDefault(true)
	if err := Configs.DatabaseInit(); err != nil {
		fmt.Println("status ", err)
	}

	Configs.DatabaseMigrate()

	r := Routers.SetupRouter()
	r.Run(":8090")
}
