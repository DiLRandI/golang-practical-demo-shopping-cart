# golang-practical-demo-shopping-cart
Simple shopping cart application to demonstrate golang and grpc

Tools used

- Docker
- Docker compose
- protoc (with golang plugin)
- Make


### Make commands

    docker-build-cart:      Build docker image of cart service.
    docker-build-item:      Build docker image of item service.
    docker-build-shipping:  Build docker image of shipping service.
    docker-build:           build all the services and docker images.

    proto-item:             Generate code from item service proto file.
    proto-shipping:         Generate code from shipping service proto file.

    env-up:                 Create containers using docker compose file.
    env-down:               Remove all the containers.
    full-cycle:             Build all the service, then create docker images for all the services, finally create containers.

    clean:                  Remove all the executable that build during *go build* operation
