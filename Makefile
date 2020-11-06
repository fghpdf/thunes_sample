PROJECT="example"

default:
	echo ${PROJECT}

unittest:
	@go test -cpu=1,2,4 -v ./...

coverage:
	@go test -race  ./... -coverprofile=coverage.txt -covermode=atomic

integration_test:
	@go test -v -tags=integration fghpdf.me/thunes_homework/internal/pkg/thunes

mock:
	@mockgen -source="internal/pkg/thunes/country/handler.go" -destination="test/mocks/thunes/country/handler.go" \
	@mockgen -source="internal/pkg/thunes/payer/handler.go" -destination="test/mocks/thunes/payer/handler.go" \
	@mockgen -source="internal/pkg/thunes/quotation/handler.go" -destination="test/mocks/thunes/quotation/handler.go" \
	@mockgen -source="internal/pkg/thunes/transaction/handler.go" -destination="test/mocks/thunes/transaction/handler.go"