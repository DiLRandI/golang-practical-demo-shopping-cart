FROM golang:alpine AS builder
WORKDIR $GOPATH/src/dilrandi/golang-practical-demo-shopping-cart
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/shipping-service .

FROM scratch
COPY --from=builder /go/bin/shipping-service /go/bin/shipping-service
EXPOSE 8080
ENTRYPOINT ["/go/bin/shipping-service"]