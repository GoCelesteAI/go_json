package main

import (
  "encoding/json"
  "fmt"
)

// Person represents a person with basic info
type Person struct {
  Name    string
  Age     int
  Email   string
  Active  bool
}

func main() {
  fmt.Println("=== JSON Marshal (Struct to JSON) ===")
  fmt.Println()

  // Create a Person struct
  person := Person{
    Name:   "Alice",
    Age:    30,
    Email:  "alice@example.com",
    Active: true,
  }

  fmt.Println("Go struct:")
  fmt.Printf("  %+v\n", person)
  fmt.Println()

  // Marshal to JSON
  jsonData, err := json.Marshal(person)
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  fmt.Println("JSON output:")
  fmt.Println(string(jsonData))
  fmt.Println()

  // Pretty print with MarshalIndent
  prettyJSON, err := json.MarshalIndent(person, "", "  ")
  if err != nil {
    fmt.Println("Error:", err)
    return
  }

  fmt.Println("Pretty JSON (MarshalIndent):")
  fmt.Println(string(prettyJSON))
}

