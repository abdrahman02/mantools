FROM golang:1.24-alpine AS build
WORKDIR /app
RUN apk add --no-cache build-base
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd

# Stage 2: production image
FROM alpine:3.19
WORKDIR /app
COPY --from=build /app/main .
COPY configs/.env ./configs/.env
EXPOSE 8080
CMD ["./main"]
