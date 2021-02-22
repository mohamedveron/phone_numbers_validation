# phone_numbers_validation
This app is developed using go-chi https://github.com/go-chi/chi router and oapi-codegen https://github.com/deepmap/oapi-codegen to generate client for apis

# Use api.yml file in the swagger ui to see the description of the rest apis


## Covered Scenarios by the Unit tests :

1- Get All Numbers categorized by country, state (valid or not valid), country code.

2- Get All Numbers filtered by one country. 

3- Get Only Valid Numbers filtered by one country.

4- Get Only Valid Numbers.

## Setup

Must have golang installed version >= 12.0.0

make file consists of 4 steps: generate, test, build, run
you can run all of them 

```bash
make all
```

# Run the unit tests:
```bash
make test
```
If you have issue with code generation in generate step you can run go test -v ./tests

# Start the http server:

```bash
make run
```

If you have issue with code generation in generate step you can copy the api/api.gen.go file in repo and run go run main.go to start the server

# Build docker image

```bash
docker build -t image-name .
```
