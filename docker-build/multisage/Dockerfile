# builder stage
FROM golang:1.13.5 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o hello ./...
# execution stage
FROM alpine:latest AS execution
COPY --from=builder /app/hello .
CMD ["./hello"]