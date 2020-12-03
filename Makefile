SRC_DIR = cmd
BUILD_DIR = build

cf8k: ${SRC_DIR}/main.go
	go build -o ${BUILD_DIR}/$@ -i $^ 

run: ${SRC_DIR}/*
	go run $^ -news ${NEWS}

test: ${SRC_DIR}/*
	go test $^

clean:
	go clean
	rm -rf ${BUILD_DIR}

.PHONY: clean
