PROJECT="example"

default:
	echo ${PROJECT}

unittest:
	@go test -cpu=1,2,4 -v ./...

coverage:
	@go test -race  ./... -coverprofile=coverage.txt -covermode=atomic

mock:
	@mockgen -source="internal/pkg/thunes/country/handler.go" -destination="test/mocks/thunes/country/handler.go" \
	@mockgen -source="internal/pkg/thunes/payer/handler.go" -destination="test/mocks/thunes/payer/handler.go"