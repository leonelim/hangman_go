BIN_NAME = hangman
SRC = ./src

all: build
build:
	go build -o $(BIN_NAME) $(SRC)
clean:
	rm $(BIN_NAME)
run: build
	./$(BIN_NAME)
help:
	@echo "targets: \
	all \
	build \
	clean \
	run \
	help" \
