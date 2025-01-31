# Goal is to generate swagger file for the API
Request and Response wrappers in `docs/create_user.go`

# To generate swagger.json file, run the following commands:

Install go-swagger:
```bash
  go install github.com/go-swagger/go-swagger/cmd/swagger@latest        
```

Generate swagger.json file:
```bash
  SWAGGER_GENERATE_EXTENSION=false swagger generate spec -o ./docs/swagger.yaml
```

