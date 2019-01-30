.PHONY: all lint test build clean show-coverage deps run-local tf-deploy

all: clean deps lint test

deps:
		@echo "Downloading dependencies..."
		@go mod download

lint:
		@echo "Run checks..."
		@go fmt $$(go list ./... | grep -v /vendor/ | grep -v /cmd)
		@go vet $$(go list ./... | grep -v /vendor/ | grep -v /cmd)

test:
		@echo "Run tests..."
		@go test -v -race $$(go list ./... | grep -v /vendor/ | grep -v /cmd) -coverprofile=coverage.out

show-coverage:
		@go tool cover -html=coverage.out

build: clean
		@mkdir -p build/
		@./build.sh

clean:
		@rm -rf build/

run-local:
		@sam local generate-event dynamodb update | sam local invoke

tf-deploy: build
		@echo "Deploying Lambda via Terraform..."
		@terraform init
		@terraform apply -auto-approve


