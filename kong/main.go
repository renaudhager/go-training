package main

import (
    "fmt"
    "kong/consumer"

)

func main() {
  fmt.Print("welcome\n")
  // Create a consumers
  consumer1, err := consumer.CreateConsumer("consumer1")
  if(err == nil){
    fmt.Printf("consumer username: %s | id: %s | created at: %s \n", consumer1.Username, consumer1.Id, consumer1.CreatedAt)
  }

  consumer2, _ := consumer.GetConsumer("consumer1")
  fmt.Print("consumer2: ", consumer2.Username, "\n")
  // consumer.ListConsumers()
  // fmt.Print("consumers: %v \n", list_consumers.Consumers)
}