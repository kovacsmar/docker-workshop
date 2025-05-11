FROM golang:1.24.2-alpine3.21 AS build

# Set destination for COPY
WORKDIR /build

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build
RUN go build -o /app/main cmd/main.go

FROM golang:1.24.2-alpine3.21


WORKDIR /app

COPY --from=build /app/main /app/main

EXPOSE 8080

# Run
CMD ["/app/main"]
