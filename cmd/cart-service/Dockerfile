FROM golang:alpine AS builder
WORKDIR $GOPATH/src/dilrandi/golang-practical-demo-shopping-cart
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/cart-service .

FROM scratch
COPY --from=builder /go/bin/cart-service /go/bin/cart-service
EXPOSE 8080
ENTRYPOINT ["/go/bin/cart-service"]