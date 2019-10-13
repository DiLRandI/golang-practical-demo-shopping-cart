.PHONY: build-cart-docker build-shopping-docker build-item-docker

docker-build-cart:
	docker build -f cmd/cart-service/Dockerfile cmd/cart-service/ -t cart-service:demo

docker-build-item:
	docker build -f cmd/item-service/Dockerfile cmd/item-service/ -t item-service:demo

docker-build-shipping:
	docker build -f cmd/shipping-service/Dockerfile cmd/shipping-service/ -t shipping-service:demo

docker-build:
	build-cart-docker build-shipping-docker build-item-docker

proto-shipping:
	protoc protos/shippingpb/shipping-service.proto --go_out=plugins=grpc:.