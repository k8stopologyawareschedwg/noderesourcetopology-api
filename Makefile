.PHONY: unit-tests
unit-tests: 
	@go test ./pkg/apis/...

.PHONY: vet
vet:
	@go vet ./pkg/...
