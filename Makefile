FOLDER = "whatever-watcher"

build:
	go build .
run:
	./$(FOLDER)
build-run:
	go build . && ./$(FOLDER)
