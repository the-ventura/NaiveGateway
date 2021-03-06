FROM node:15.5.0 as frontend-builder
WORKDIR /usr/src/app
COPY ./web/frontend/package.json package.json
RUN yarn install
COPY ./web/frontend .
RUN yarn build

FROM golang:1.16 AS builder
WORKDIR /go/src/github.com/the-ventura/naivegateway/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o naivegateway cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/the-ventura/naivegateway/naivegateway .
COPY --from=frontend-builder /usr/src/app/build ./web/frontend/build
COPY MIGRATION_VERSION MIGRATION_VERSION
COPY migrations migrations

CMD ["./naivegateway", "api"]
