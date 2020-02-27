go get github.com/golang/protobuf/protoc-gen-go

protoc --proto_path=./proto --go_out=./src_gen score_info.proto