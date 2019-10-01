package main

import (
	"fmt"

	"github.com/YangChinFu/config-env/config"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := config.New()
	fmt.Printf("%+v", config)
}
