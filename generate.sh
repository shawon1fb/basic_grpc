
#protoc greet\\greetpb\\greet.proto --go_out=grpc:pb
#protoc --proto_path=proto/ proto/greet.proto --go_out=plugins=grpc:pb/
#protoc --proto_path=greet/greetpb/ greet/greetpb/greet.proto --go_out=plugins=grpc:greet/greetpb --go-grpc_out=:plugins=grpc:greet/greetpb

#protoc -I protos/greetpb/ protos/greetpb/greet.proto --dart_out=grpc:lib/src/generated


#protoc --go-grpc_out=pb proto/*.proto --go_out=:pb --go-grpc_out=:pb

protoc --proto_path=greet/greetpb/ greet/greetpb/greet.proto --go-grpc_out=greet/greetpb