##@ General

.PHONY: help
help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target>\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  %-15s %s\n", $$1, $$2 } /^##@/ { printf "\n%s\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: vet
vet: ## Run go vet against the code
	go vet ./...

.PHONY: vuln
vuln: govulncheck ## Run govulncheck against the code
	$(GOVULNCHECK) ./...

.PHONY: lint
lint: staticcheck ## Run staticcheck against the code
	$(STATICCHECK) ./...

.PHONY: test
test: lint vet vuln ## Run tests
	go test -v -covermode=count -coverprofile=cover.out ./...

.PHONY: build
build: lint vet vuln ## Build git-credential-op
	go build -o $(LOCALBIN)/git-credential-op cmd/git-credential-op/main.go

##@ Build Dependencies

## Location to install dependencies to
LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

## Tool binaries
GOVULNCHECK ?= $(LOCALBIN)/govulncheck
STATICCHECK ?= $(LOCALBIN)/staticcheck

## Tool versions
GOVULNCHECK_VERSION ?= v1.0.1
STATICCHECK_VERSION ?= 2023.1.6

.PHONY: staticcheck
staticcheck: $(STATICCHECK) ## Download staticcheck locally if needed
$(STATICCHECK): $(LOCALBIN)
	test -s $(LOCALBIN)/staticcheck && $(LOCALBIN)/staticcheck --version | grep -q $(STATICCHECK_VERSION) || \
	GOBIN=$(LOCALBIN) go install honnef.co/go/tools/cmd/staticcheck@$(STATICCHECK_VERSION)

.PHONY: govulncheck
govulncheck: $(GOVULNCHECK)
$(GOVULNCHECK): $(LOCALBIN)
	test -s $(LOCALBIN)/govulncheck && $(LOCALBIN)/govulncheck --version | grep -q $(GOVULNCHECK_VERSION) || \
	GOBIN=$(LOCALBIN) go install golang.org/x/vuln/cmd/govulncheck@$(GOVULNCHECK_VERSION)
