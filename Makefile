libnostr.so: $(shell find . --name "*.go")
	go build -buildmode=c-shared -o libnostr.so

install: libnostr.so
	mv libnostr.so ~/.purple/plugins/
