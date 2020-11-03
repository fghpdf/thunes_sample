mock:
	mockgen -source="internal/pkg/thunes/country/handler.go" -destination="test/mocks/thunes/country/handler.go"
	mockgen -source="internal/pkg/thunes/payer/handler.go" -destination="test/mocks/thunes/payer/handler.go"