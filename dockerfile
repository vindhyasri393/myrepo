FROM golang:1.24.1 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o helloworldapp

FROM debian:bookworm-slim
WORKDIR /root/
COPY --from=builder /app/helloworldapp .
RUN chmod +x helloworldapp
CMD ["./helloworldapp"]
