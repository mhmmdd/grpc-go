
In order to generate grp.go file run these commands.
`go get -u google.golang.org/grpc`
`go get -u github.com/golang/protobuf/protoc-gen-go`

`protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.`
`protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.`
`protoc blog/blogpb/blog.proto --go_out=plugins=grpc:.`

Error Codes
* https://grpc.io/docs/guides/error/
* http://avi.im/grpc-errors/

Deadlines
* https://grpc.io/blog/deadlines/

SSL Certificate
* https://grpc.io/docs/guides/auth/

## Evans 
1. Download the evans: https://github.com/ktr0731/evans/releases
1. Add this line to server.go
    `// Register reflection service on gRPC server.
    reflection.Register(s)`
1. Add the environment path and run this command: 
   `evans -p 50051 -r`
1. Some evans commands: 
    `show package
    show service
    show message
    desc MESSAGE_NAME`

    `show service
    service SERVICE_NAME
    call RPC_NAME`
    
# Blog
1. Install [MinGW](https://sourceforge.net/projects/tdm-gcc/) and add it the environment path (C:\TDM-GCC-64\bin). That fixes this [issue.](https://stackoverflow.com/questions/43580131/exec-gcc-executable-file-not-found-in-path-when-trying-go-build)
1. Install MongoDB to Docker
`docker run -d -p 27017:27017 --name mongo \
    -e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
    -e MONGO_INITDB_ROOT_PASSWORD=secret \
    mongo`