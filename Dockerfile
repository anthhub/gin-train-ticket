FROM golang:1.15-alpine AS builder
ADD main.go /go/src/hello/main.go
WORKDIR /go/src/hello
RUN go build main.go

FROM alpine:3.13
EXPOSE 8000
COPY --from=builder /go/src/hello/main /main
CMD ["/main"]