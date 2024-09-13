# local server
.PHONY: run-server
run-server:
	source ./.env && cd ./src && go run ./main.go -f ./etc/definition.yaml
