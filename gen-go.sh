if [ ! -d "./gen/go" ]; then
	mkdir -p "./gen/go"
fi

protoc --proto_path=./protos --go-grpc_out=./gen/go ./protos/*/*.proto

