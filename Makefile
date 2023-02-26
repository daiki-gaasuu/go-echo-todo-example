OAPI = $(shell which oapi-codegen)
OPENAPI = $(shell which openapi-generator)

gen-backend:
	$(OAPI) -old-config-style -generate server -package openapi -o ./server/api/generated/server.gen.go ./server/api/openapi.yaml
	$(OAPI) -old-config-style -generate types -package openapi -o ./server/api/generated/types.gen.go ./server/api/openapi.yaml
	$(OAPI) -old-config-style -generate spec -package openapi -o ./server/api/generated/spec.gen.go ./server/api/openapi.yaml
	go generate ./...

gen-frontend:
	${OPENAPI} generate -i server/api/openapi.yaml -g typescript-axios -o client/src/api/api-client