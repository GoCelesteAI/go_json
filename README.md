# Working with JSON in Go

Demo code from **Go Tutorial Lesson 23: Working with JSON**

Watch the full tutorial: [YouTube Link]

## Topics Covered

1. **json.Marshal** - Convert Go structs to JSON bytes
2. **json.Unmarshal** - Parse JSON into Go structs
3. **Struct Tags** - Control JSON field names and behavior
4. **Nested Structures** - Handle nested JSON objects

## Files

| File | Description |
|------|-------------|
| `json_marshal.go` | Convert struct to JSON with Marshal and MarshalIndent |
| `json_unmarshal.go` | Parse JSON string into Go struct |
| `json_tags.go` | Struct tags for custom names, omitempty, and exclusion |
| `json_nested.go` | Nested structs for complex JSON structures |

## Running the Examples

```bash
# Marshal - Struct to JSON
go run json_marshal.go

# Unmarshal - JSON to Struct
go run json_unmarshal.go

# Struct Tags
go run json_tags.go

# Nested Structures
go run json_nested.go
```

## Key Concepts

### json.Marshal
```go
jsonData, err := json.Marshal(person)
prettyJSON, err := json.MarshalIndent(person, "", "  ")
```

### json.Unmarshal
```go
var product Product
err := json.Unmarshal([]byte(jsonString), &product)
```

### Struct Tags
```go
type User struct {
    ID       int    `json:"id"`              // Custom name
    Email    string `json:"email,omitempty"` // Skip if empty
    Password string `json:"-"`               // Exclude from JSON
}
```

### Nested Structures
```go
type Company struct {
    Name    string  `json:"name"`
    Address Address `json:"address"` // Nested struct
}
```

## Common Tag Options

| Tag | Description |
|-----|-------------|
| `json:"name"` | Use "name" as the JSON key |
| `json:"name,omitempty"` | Skip field if zero value |
| `json:"-"` | Exclude field from JSON |

## Channel

Subscribe for more Go tutorials: [GoCeleste AI](https://youtube.com/@GoCelesteAI)
