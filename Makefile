.PHONY: openapi_http
openapi_http:
	@./scripts/openapi-http.sh ports ports

.PHONY: proto
proto:
	@./scripts/proto.sh monNlpService