# Go Gin Chunk API

This project is a simple API built with Gin in Go that takes a number input, divides it into multiple chunks, and processes those chunks concurrently using goroutines.

## Project Structure

```
go-gin-chunk-api
├── main
│   └── main.go          # Entry point of the application
├── repository
│   ├── handler
│   │   └── api.go      # Handles incoming requests
│   ├── service
│   │   └── chunk.go     # Processes chunks concurrently
│   └── model
│       └── request.go   # Defines request structure
├── go.mod               # Module definition
└── README.md            # Project documentation
```

