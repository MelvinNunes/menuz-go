# STARTUP (VERSION 1.0)

Hey there,

## GETTING STARTED

To get started with the VOLET API you need to have Docker and Docker Compose installed (if you're only trying to run the app) and GoLang installed (if youre trying to run it by yourself).

- Docker Version that Was Used: 4.22.1 (118664)

- Or you can run it the old good way setting up a mysql database and GoLang enviroment. 😅😅

Firstly, you need to setup all the environment variables, create a new .env file and configure it like the .env.example file that you can find in the root of the project.

After setting up all environment variables, just run "docker compose up -d" and then wait for docker compose to start all services (containers) and the you can check "localhost:8000/v2/health" (the standard port number is 8000) to check if the app is running!

- If youre willing to run it without docker, just set up the database in .env and install go lang (v1.21.3 or more) and the use the command "go run cmd/api/main.go" in the root of the project.

You'll check the Fiber terminal UI to confirm that the app is running!

## INTERNATIONALIZATION

This API allows you to choose between two different languages (Portuguese - pt as Default and English - en) while doing requests, just add the `Accept-Language` header with `en` for English or `pt` for Portuguese.

## DOCUMENTATION

You can also use the following url "localhost:8000/docs" to see the documentation (we used swagger for the API documentation)

Install swaggo (https://github.com/swaggo/swag?tab=readme-ov-file#supported-web-frameworks) so you can updated the documentation. Using "go install github.com/swaggo/swag/cmd/swag@latest"

To updated the documentation just run:

- swag init -d internal/interfaces -g ../../cmd/api/main.go

## MIGRATION

To create a new migration you can use the following command:

- migrate create -ext sql -dir internal/infrastructure/database/migrations -seq migration_name

## TESTING

For testing purposes, we use the standard golang test library, you can run "go test ./..." to run all tests

## DEPLOYMENT
