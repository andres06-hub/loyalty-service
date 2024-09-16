# local server
.PHONY: run-server
run-server:
	source ./.env && cd ./src && go run ./main.go -f ./etc/definition.yaml

swag:
	swag fmt && swag init --output ./src/docs

# Variables
COVERAGE_FILE := coverage.out
COVERAGE_HTML := coverage.html

.PHONY: test coverage coverage-html clean

test:
	go test -v -test.short -p=1 -run "^*__Unit" ./src/test/...

coverage:
	go test -coverprofile=$(COVERAGE_FILE) ./...
	go tool cover -func=$(COVERAGE_FILE)

coverage-html: coverage
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "Coverage report generated: $(COVERAGE_HTML)"

clean:
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)

all: test coverage