package main

import (
	"api-gateway/src/pkgs/gorm"
	"fmt"
)

func main() {
	gorm.CreateDbInstance()

	fmt.Println(("Loaded api-gateway!"))
}
