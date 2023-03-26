# Build stage
FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

# Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main .
COPY --from=build /app/.env .

EXPOSE 8080

CMD ["./main"]