# export PATH=$PATH:$HOME/go/bin
# export PATH=$PATH:/usr/local/go/bin
protoc --go-grpc_out=./internal/protorender ./proto/*.proto