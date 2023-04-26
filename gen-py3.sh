if [ ! -d "./gen/py3" ]; then
	mkdir -p "./gen/py3"
fi

python3 -m grpc_tools.protoc --proto_path=. -I./protos --python_out=./gen/py3 --grpc_python_out=./gen/py3 ./protos/*/*.proto
echo "proto file create success ..."