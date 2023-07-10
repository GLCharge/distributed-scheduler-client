gen-client:
	oapi-codegen -generate types,client ./api/openapi.yaml > scheduler-client.gen.go