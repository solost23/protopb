if [ ! -d "./gen/go" ]; then
	mkdir -p "./gen/go"
fi

protoc --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative ./protos/*/*.proto
