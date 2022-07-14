#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git

WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN ls /go/bin/
RUN go build -o /go/bin/app

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=minitest Version=0.0.1
EXPOSE 8081