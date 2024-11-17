package conf

import (
	"fmt"
	"os"
)

func GetVar(name string) string {
	value := os.Getenv(name)
	fmt.Println(name + " = " + value)
	return value
}
