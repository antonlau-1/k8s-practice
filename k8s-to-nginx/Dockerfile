FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN go build -o k8s_to_nginx .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/k8s_to_nginx .

EXPOSE 8080 

CMD ["./k8s_to_nginx"]