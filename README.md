# grpc-mongo
Crud to mongo with grpc

install mongo
1. go get go.mongodb.org/mongo-driver/mongo


requirements:
1. mongodb (docker)
2. mongo client (compass or robo-3T)
3. mongo-go-driver (sqlalchemy for go-mongodb)
    docs
    a. (https://docs.mongodb.com/drivers/go/)
    b. https://www.mongodb.com/blog/search/golang%20quickstart


export-import notes
1. go only export func and variable that is begun with Capital


Folder Description
1. connection: it consists of connection to mongo db, it has function that return client and db that can be reused
2. mongo101: contains basic operation of mongodb in go like, create, read, update, and delete
4. model: contain object mapping to mongo
5. proto: is where protobuf grpc live
6. service: include all function that can be called by client (remote function)
7. server: this location is entrypoint of this app


How to run:
1. look at .proto file in proto folder to gain insight at beginning
2. generate all .proto with generate.sh
3. see all service that can be called remotely by looking at service folder
4. tidy up go module by running "go mod tidy"
5. run with make file, for server "make starts", for client "make startc"