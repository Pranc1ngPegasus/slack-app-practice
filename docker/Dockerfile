FROM golang:1.16 as builder
WORKDIR /go/src/github.com/Pranc1ngPegasus/slack-api-practice
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/Pranc1ngPegasus/slack-api-practice/app .
CMD ["./app"]
