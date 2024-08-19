From golang;1.22.5-alpine3.18 as builder

Run mkdir app

Copy . /app

Workdir /app
Run go build -o main main.go
From alpine:3.18
Workdir /app
Copy --from=builder /app .
CMD ["/app/main"]