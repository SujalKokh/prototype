package main

import (
  "fmt"
  "log"
  "github.com/joho/godotenv"
  "prototype/config"
)

func main(){
	var err error

  err = godotenv.Load()
  if err != nil {
    log.Fatalf("Error getting env, %v", err)
  } else {
    fmt.Println("Environment variables fetched ")
  }
  
  config.InitServer()
}