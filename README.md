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
1. connection:
it consists of connection to mongo db, it has function that return client and db that can be reused
2. mongo101
contains basic operation of mongodb in go like, create, read, update, and delete
4. model contain object mapping to mongo
5. proto is where protobuf grpc live