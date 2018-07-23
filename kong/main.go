package main

import (
    "fmt"
    "kong/consumer"

)

func main() {
  fmt.Print("welcome\n")
  // Create a consumers
  consumer.CreateConsumer("go_consumer")
  // fmt.Print("consumer: ", consumer_1.Username)

  consumer.ListConsumers()
  // fmt.Print("consumers: %v \n", list_consumers.Consumers)
}