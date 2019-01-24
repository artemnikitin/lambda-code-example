.PHONY: all lint test build clean

all: clean lint test

lint:
		@echo "Run checks..."
		@go fmt $$(go list ./... | grep -v /vendor/ | grep -v /cmd)
		@go vet $$(go list ./... | grep -v /vendor/ | grep -v /cmd)

test:
		@echo "Run tests..."
		@go test -v -race $$(go list ./... | grep -v /vendor/ | grep -v /cmd)

build: clean
		@mkdir -p build/
		./build.sh
		@find build -type f ! -name '*.zip' -delete

clean:
		@rm -rf build/
