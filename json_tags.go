package main

import (
  "encoding/json"
  "fmt"
)

// User demonstrates struct tags for JSON
type User struct {
  ID        int    `json:"id"`
  Username  string `json:"username"`
  Email     string `json:"email,omitempty"`
  Password  string `json:"-"`
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
  Age       int    `json:"age,omitempty"`
}

func main() {
  fmt.Println("=== JSON Struct Tags ===")
  fmt.Println()

  // User with all fields
  user1 := User{
    ID:        1,
    Username:  "johndoe",
    Email:     "john@example.com",
    Password:  "secret123",
    FirstName: "John",
    LastName:  "Doe",
    Age:       25,
  }

  fmt.Println("User with all fields:")
  printJSON(user1)

  // User with empty optional fields
  user2 := User{
    ID:        2,
    Username:  "janedoe",
    Password:  "hidden456",
    FirstName: "Jane",
    LastName:  "Doe",
    // Email and Age are empty (zero values)
  }

  fmt.Println("User with empty optional fields:")
  printJSON(user2)

  fmt.Println("Tag meanings:")
  fmt.Println("  json:\\\"name\\\"     - Custom JSON key name")
  fmt.Println("  json:\\\"-\\\"        - Exclude from JSON")
  fmt.Println("  json:\\\"omitempty\\\"- Skip if zero value")
}

func printJSON(v any) {
  data, _ := json.MarshalIndent(v, "", "  ")
  fmt.Println(string(data))
  fmt.Println()
}

