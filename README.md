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

## Getting Started

### Prerequisites

- Go 1.16 or later
- Gin framework

### Installation

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/go-gin-chunk-api.git
   cd go-gin-chunk-api
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

### Running the API

To run the API, execute the following command:

```
go run main/main.go
```

The server will start on `http://localhost:8080`.

### API Endpoints

#### POST /:number

This endpoint accepts a number as a path parameter and divides it into chunks for processing.

**Example Request:**

```bash
curl -X POST http://localhost:8080/1000
```

**Response:**

A JSON response indicating the status of the processing.

### Example

To test the API, you can use `curl`:

```bash
curl -X POST http://localhost:8080/chunk -H "Content-Type: application/json" -d '{"number": 100}'
```

### Contributing

Feel free to submit issues or pull requests for improvements or bug fixes.

### License

This project is licensed under the MIT License.