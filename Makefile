
gen:
	protoc --proto_path=greet/greetpb/ greet/greetpb/greet.proto --go-grpc_out=greet/greetpb --go_out=greet/greetpb --go-grpc_opt=require_unimplemented_servers=false

server:
	go run .\greet_server\server.go

client:
	go run .\greet_client\client.go


#gen2:
#	protoc --proto_path=greet/greetpb/ greet/greetpb/greet.proto --go-grpc_out=greet/greetpb --go_out=greet/greetpb --go-grpc_opt=require_unimplemented_servers=false --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
