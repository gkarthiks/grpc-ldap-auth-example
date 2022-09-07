DIST_DIR=$(CURDIR)/dist

gen:
	protoc -I=proto/ --go_out=proto/ --go-grpc_out=proto/ proto/grpc-ldap.proto

serve:
	go run main.go

cli:
	go build -o ${DIST_DIR}/greet greet/main.go

clean:
	rm -f ${DIST_DIR}/greet
