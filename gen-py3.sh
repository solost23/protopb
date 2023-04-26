if [ ! -d "./gen/py3" ]; then
	mkdir -p "./gen/py3"
fi

# --proto_path=. 可替换成-I.
python3 -m grpc_tools.protoc --proto_path=. --python_out=./gen/py3 --grpc_python_out=./gen/py3 ./protos/*/*.proto
echo "proto file create success ..."