# API Development in Go
Master the concept of microservices API development using the Hexagonal Architecture in Go.

## Task 1 - Build a Time API

Task 1 implementation, building an API that returns the time in the UTC timezone or the time in different timezones passed by parameter in the request

### Project Structure

The project structure is as follows:

```
api-some-time-inc/
├── go.mod
├── go.sum
├── main.go
├── api/
│   ├── api.go
│   └── handlers.go
```

### Files

- **main.go**: Application entry point.
- **api/api.go**: Application configuration.
- **api/handlers.go**: Handlers for HTTP requests.

### Running the Application

1. Navigate to the project directory and run `go run main.go`.
2. The API will be available at `http://localhost:8080/api/time`
3. It is possible to add parameters to the request using the 'tz' tag

### Testing the API

Use tools like `curl` or `Postman` to test the `/api/time` endpoint:

UTC time
```sh
curl http://localhost:8080/api/time
```

Time in another timezones
```sh
curl http://localhost:8000/api/time?tz=America/New_York,Asia/Kolkata
```

### Additional Resources

- [Database Time Zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
