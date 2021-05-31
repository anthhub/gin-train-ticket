FROM golang:1.15-alpine AS builder
ADD . /go/src/myapp
WORKDIR /go/src/myapp
RUN go build ./main.go

FROM alpine:3.13
EXPOSE 8000
COPY --from=builder /go/src/myapp/main /main
COPY static ./static
CMD ["/main"]



