.PHONY: build-cart-docker build-shopping-docker build-item-docker

build-cart-docker:
	docker build -f cmd/cart-service/Dockerfile cmd/cart-service/ -t cart-service:demo

build-item-docker:
	docker build -f cmd/item-service/Dockerfile cmd/item-service/ -t item-service:demo

build-shipping-docker:
	docker build -f cmd/shipping-service/Dockerfile cmd/shipping-service/ -t shipping-service:demo

build-docker:
	build-cart-docker build-shipping-docker build-item-docker