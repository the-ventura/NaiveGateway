FROM golang:1.16 AS builder
WORKDIR /go/src/github.com/the-ventura/naivegateway/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o naivegateway cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/the-ventura/naivegateway/naivegateway .
CMD ["./naivegateway", "api"]
