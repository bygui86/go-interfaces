
# Go testing

Project to understand how to write testing in Golang

`WORK IN PROGRESS`

## Test
```
go test ./...
```

## Code coverage

### Test
```
go test -cover ./...
```

### Show results
```
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out
```

---

## Links

- https://medium.com/swlh/using-go-interfaces-for-testable-code-d2e11b02dea
- https://blog.golang.org/cover
