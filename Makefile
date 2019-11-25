.PHONY: docker-build-cart docker-build-item docker-build-shipping env-up

PROJECT_FOLDER=/go/src/github.com/dilrandi/golang-practical-demo-shopping-cart
GO_IMG=golang:alpine
GO_VOL=-v $(GOPATH)/src:/go/src
GO_ENV=-e CGO_ENABLED=0 -e GOOS=linux
GO_WD=-w $(PROJECT_FOLDER)
GO_CMD=docker run --rm -i $(GO_VOL) $(GO_ENV) $(GO_WD) $(GO_IMG) go build -a


env-up:
	cd ./cmd/compose && docker-compose up -d

env-down:
	cd ./cmd/compose && docker-compose down -v

docker-build-cart:
	$(GO_CMD) -o cmd/cart-service/bin/cart-service $(PROJECT_FOLDER)/cmd/cart-service/main.go
	docker build -f cmd/cart-service/Dockerfile cmd/cart-service/ -t cart-service:demo

docker-build-item:
	$(GO_CMD) -o cmd/item-service/bin/item-service $(PROJECT_FOLDER)/cmd/item-service/main.go
	docker build -f cmd/item-service/Dockerfile cmd/item-service/ -t item-service:demo

docker-build-shipping:
	$(GO_CMD) -o cmd/shipping-service/bin/shipping-service $(PROJECT_FOLDER)/cmd/shipping-service/main.go
	docker build -f cmd/shipping-service/Dockerfile cmd/shipping-service/ -t shipping-service:demo

docker-build:env-up docker-build-cart docker-build-item

proto-shipping:
	protoc protos/shippingpb/shipping-service.proto --go_out=plugins=grpc:.

clean: 
	rm -rf cmd/item-service/bin cmd/cart-service/bin cmd/shipping-service/bin