# Stage 1: Build
FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /main .

# Stage 2: Final image
FROM golang:1.24.1

WORKDIR /app

COPY --from=builder /main .

EXPOSE 8080

CMD ["./main"]
