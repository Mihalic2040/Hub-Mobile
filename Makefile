build_core:
	GO111MODULE=on gomobile bind -target android -androidapi 31 -o ./android/app/libs/deamon.aar github.com/Mihalic2040/Hub-Mobile

init:
	gomobile init

test:
	cd examples && go run main.go

grpc:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protocols/bridge.proto


	protoc --java_out=lite:android/app/src/main/java --grpc-java_out=lite:android/app/src/main/java protocols/bridge.proto
