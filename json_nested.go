package main

import (
  "encoding/json"
  "fmt"
)

// Address represents a physical address
type Address struct {
  Street  string `json:"street"`
  City    string `json:"city"`
  Country string `json:"country"`
  ZipCode string `json:"zip_code"`
}

// Company represents a business entity
type Company struct {
  Name      string   `json:"name"`
  Address   Address  `json:"address"`
  Employees []string `json:"employees"`
  Founded   int      `json:"founded"`
}

func main() {
  fmt.Println("=== Nested JSON Structures ===")
  fmt.Println()

  // Create nested structure
  company := Company{
    Name: "Tech Corp",
    Address: Address{
      Street:  "123 Main Street",
      City:    "San Francisco",
      Country: "USA",
      ZipCode: "94102",
    },
    Employees: []string{"Alice", "Bob", "Charlie"},
    Founded:   2020,
  }

  // Marshal to JSON
  jsonData, _ := json.MarshalIndent(company, "", "  ")
  fmt.Println("Nested JSON output:")
  fmt.Println(string(jsonData))
  fmt.Println()

  // Unmarshal nested JSON
  jsonInput := `{
    "name": "Startup Inc",
    "address": {
      "street": "456 Tech Ave",
      "city": "Austin",
      "country": "USA",
      "zip_code": "78701"
    },
    "employees": ["Dave", "Eve"],
    "founded": 2023
  }`

  var parsed Company
  json.Unmarshal([]byte(jsonInput), &parsed)

  fmt.Println("Parsed nested JSON:")
  fmt.Printf("  Company: %s (Founded: %d)\n", parsed.Name, parsed.Founded)
  fmt.Printf("  Location: %s, %s\n", parsed.Address.City, parsed.Address.Country)
  fmt.Printf("  Team: %v\n", parsed.Employees)
}

