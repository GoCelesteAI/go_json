package main

import (
  "encoding/json"
  "fmt"
)

// Product represents an item in a store
type Product struct {
  ID    int
  Name  string
  Price float64
  Stock int
}

func main() {
  fmt.Println("=== JSON Unmarshal (JSON to Struct) ===")
  fmt.Println()

  // JSON data as a string
  jsonString := `{
    "ID": 101,
    "Name": "Laptop",
    "Price": 999.99,
    "Stock": 50
  }`

  fmt.Println("JSON input:")
  fmt.Println(jsonString)
  fmt.Println()

  // Create empty struct to hold data
  var product Product

  // Unmarshal JSON into struct
  err := json.Unmarshal([]byte(jsonString), &product)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  fmt.Println("Go struct:")
  fmt.Printf("  ID:    %d\n", product.ID)
  fmt.Printf("  Name:  %s\n", product.Name)
  fmt.Printf("  Price: $%.2f\n", product.Price)
  fmt.Printf("  Stock: %d units\n", product.Stock)
  fmt.Println()

  // Calculate total value
  totalValue := product.Price * float64(product.Stock)
  fmt.Printf("Total inventory value: $%.2f\n", totalValue)
}

