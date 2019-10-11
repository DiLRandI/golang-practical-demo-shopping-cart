FROM golang:alpine AS builder
WORKDIR $GOPATH/src/dilrandi/golang-practical-demo-shopping-cart
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/item-service .

FROM scratch
COPY --from=builder /go/bin/item-service /go/bin/item-service
EXPOSE 8080
ENTRYPOINT ["/go/bin/item-service"]