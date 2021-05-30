FROM golang:1.15-alpine AS builder
ADD . /go/src/myapp
WORKDIR /go/src/myapp
RUN ls
RUN go build ./main.go

FROM alpine:3.13
EXPOSE 8000
COPY --from=builder /go/src/myapp/main /main
CMD ["/main"]

# FROM golang:1.15-alpine AS builder
# ADD . /go/src/myapp
# WORKDIR /go/src/myapp
# # RUN go get myapp
# RUN go install ./gin-train-ticket/...

# FROM alpine:3.13
# EXPOSE 8000
# COPY --from=builder /go/bin/gin-train-ticket /bin/gin-train-ticket
# CMD ["/bin/gin-train-ticket"]


