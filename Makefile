#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export CGO_ENABLED=0
export GO111MODULE=on

.PHONY: hp
hp: ## show list command 
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := hp

###

.which-go:
	@which go > /dev/null || (echo "install go from https://golang.org/dl/" & exit 1)

.PHONY: format
format: .which-go ## Formats Go files
	gofmt -s -w $(ROOT)

.which-lint:
	@which golangci-lint > /dev/null || (echo "install golangci-lint from https://github.com/golangci/golangci-lint" & exit 1)

.PHONY: lint
lint: .which-lint ## Checks code with Golang CI Lint
	golangci-lint run

.PHONY: clean
clean: # run make format and make lint
	gofmt -s -w $(ROOT)
	golangci-lint run
	
.PHONY: com
com: ## git commit
	git add .
	git commit -a

.PHONY: test
test: ## run test
	go test ./... -v
