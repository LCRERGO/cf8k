SRC_DIR = cmd
BUILD_DIR = build

cf8k: ${SRC_DIR}/*.go
	go build -o ${BUILD_DIR}/$@ -i $^ 

run: ${SRC_DIR}/*.go
	go run $^ -news ${NEWS}

doc: ${SRC_DIR}
	go doc -all $^

clean:
	go clean
	rm -rf ${BUILD_DIR}

.PHONY: clean
