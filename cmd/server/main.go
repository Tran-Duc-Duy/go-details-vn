package main

import (
	"go-tip/internal/routers"
)

func main() {
  r := routers.NewRouter()

  r.Run() 
}
