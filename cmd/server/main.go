package main

import (
	"my-ecom.com/api/internal/routers"
)

func main() {
	r := routers.NewRouter()

	r.Run(":8000") // default port 8080
}
