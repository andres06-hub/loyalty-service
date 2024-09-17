# local server
.PHONY: run-server
run-server:
	source ./.env && cd ./src && go run ./main.go -f ./etc/definition.yaml

swag:
	swag fmt && swag init -g main.go -d ./src -o src/docs

# Variables
COVERAGE_FILE := coverage.out
COVERAGE_HTML := coverage.html

.PHONY: test coverage coverage-html clean

test:
	make coverage
	make coverage-html
	go test -v -test.short -p=1 -run "^*__Unit" ./src/test/...
	@coverage_path="file://$(PWD)/coverage/coverage.html"; \
  echo $$coverage_path

coverage:
	go test -coverprofile=./coverage/$(COVERAGE_FILE) ./...
	go tool cover -func=./coverage/$(COVERAGE_FILE)

coverage-html: coverage
	go tool cover -html=./coverage/$(COVERAGE_FILE) -o ./coverage/$(COVERAGE_HTML)
	@echo "Coverage report generated: $(COVERAGE_HTML)"

clean:
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)

test-all: test coverage