.PHONY: docker-build-cart docker-build-item docker-build-shipping env-up

PROJECT_FOLDER=/go/src/github.com/dilrandi/golang-practical-demo-shopping-cart

GO_IMG=golang:alpine
GRPC_IMG=deleema1/grpc-tool-golang
GODOG_IMG=deleema1/godog

PWD=$(shell pwd)
GO_VOL=-v $(GOPATH)/src:/go/src
GRPC_VOL=-v $(PWD):/opt/mnt

GO_ENV=-e CGO_ENABLED=0 -e GOOS=linux
GO_WD=-w $(PROJECT_FOLDER)
GODOG_WD=-w $(PROJECT_FOLDER)/integration_tests

GO_CMD=docker run --rm -i $(GO_VOL) $(GO_ENV) $(GO_WD) $(GO_IMG) go build -a
GRPC_CMD=docker run --rm -i $(GRPC_VOL) $(GRPC_IMG)
GODOG_CMD=docker run --rm -i $(GO_VOL) $(GODOG_WD) $(GODOG_IMG) godog

docker-build-cart:
	$(GO_CMD) -o cmd/cart-service/bin/cart-service $(PROJECT_FOLDER)/cmd/cart-service/main.go
	docker build -f cmd/cart-service/Dockerfile cmd/cart-service/ -t cart-service:demo

docker-build-item:
	$(GO_CMD) -o cmd/item-service/bin/item-service $(PROJECT_FOLDER)/cmd/item-service/main.go
	docker build -f cmd/item-service/Dockerfile cmd/item-service/ -t item-service:demo

docker-build-shipping:
	$(GO_CMD) -o cmd/shipping-service/bin/shipping-service $(PROJECT_FOLDER)/cmd/shipping-service/main.go
	docker build -f cmd/shipping-service/Dockerfile cmd/shipping-service/ -t shipping-service:demo

docker-build: docker-build-cart docker-build-item #docker-build-shipping

proto-item:
	$(GRPC_CMD) protos/itempb/item-service.proto --go_out=plugins=grpc:.

proto-shipping:
	$(GRPC_CMD) protos/shippingpb/shipping-service.proto --go_out=plugins=grpc:.
	
env-up:
	cd ./compose && docker-compose up -d

env-down:
	cd ./compose && docker-compose down -v

full-cycle : docker-build env-up

clean: 
	rm -rf cmd/item-service/bin cmd/cart-service/bin cmd/shipping-service/bin