tidy:
	go mod tidy; go mod vendor;

gen:
	protoc --go_out=:api/v1/proto/pb --go-grpc_out=:api/v1/proto/pb ./api/v1/proto/definition/*.proto

run:
	. ./dev.rc && (cd cmd/api; go build . && ./api; cd ../..);
