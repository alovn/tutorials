all:
	openssl genrsa -out server.key 2048
	openssl req -new -x509 -key server.key -out server.pem -days 3650 -subj "/C=CN/ST=ZJ/L=HZ/O=grpc/OU=dev/CN=localhost:8888/emailAddress=me@grpc.io"
