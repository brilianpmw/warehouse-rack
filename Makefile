build:
	go build -v -o bin/warehouse_rack main.go

run:
	./bin/warehouse_rack

BINARY=./bin/warehouse_rack

run-file:
	@echo "Running $(BINARY) with $(FILE)"
	$(BINARY) $(FILE)