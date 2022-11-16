
doc:
	godoc --http :8080

doc_swagger:
	swag init

install_swagger:
	go install github.com/swaggo/swag/cmd/swag@latest

test:
	go test

run_local: doc_swagger 
	export PORT=8080 && go run ./main.go

new_version:
	./scripts/update_version.sh

config_git_hooks:
	git config core.hooksPath .githooks
