gen:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative  processor_message.proto memory_message.proto storage_message.proto keyboard_message.proto screen_message.proto labtop_message.proto

clean:
	rm pb/*go 

run:
	go run main.
	
test:
	go test -cover -race ./...