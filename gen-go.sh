if [ ! -d "./gen/go" ]; then
	mkdir -p "./gen/go"
fi

protoc --proto_path=. --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative ./*/*.proto
echo "proto file create success ..."