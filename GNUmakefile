default: testacc

.PHONY: testacc test

testacc:
	@TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

test:
	@go test ./...

.PHONY: build

build:
	@go build -o dist/terraform-provider-drone

.PHONY: format lint

format:
	@goimports -w -e -local github.com/FriendsOfDrone internal . && \
	gofumpt -w {} && \
	gofmt -w -s {}

lint:
	@golangci-lint run --deadline=5m0s --out-format=line-number -exclude-use-default=false ./...

.PHONY: check-upgrades upgrade

check-upgrades:
	@go list -mod=readonly -u \
		-f "{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}" -m all

upgrade:
	@go get -u ./... && go mod tidy
