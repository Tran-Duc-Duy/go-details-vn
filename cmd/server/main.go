package main

import (
	"go-tip/internal/routers"
	"log"
)

func main() {
  r := routers.NewRouter()


  for _, route := range r.Routes() {
    log.Printf("Method: %s, Path: %s", route.Method, route.Path)
  }
  r.Run() 
}
