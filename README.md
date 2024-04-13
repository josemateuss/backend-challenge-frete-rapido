# Backend Challenge - Frete Rápido

This is my project for the Backend Developer job application at Frete Rápido.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Prerequisites

- Docker

### Execution

1. Clone the repository:

```bash
git clone https://github.com/josemateuss/backend-challenge-frete-rapido.git
```

2. Navigate to the project directory:

```bash
cd backend-challenge-frete-rapido
```

3. Run Docker Compose:

```bash
docker-compose up --build
```

The application will be running at `localhost:3000`.

## API Endpoints

Here are the endpoints required by the challenge:

- POST: /quote:

```bash
curl --location 'localhost:3000/frete-rapido/api/v1/quote' \
--header 'Content-Type: application/json' \
--data '{"recipient":{"address":{"zipcode":"73340030"}},"volumes":[{"category":7,"amount":1,"unitary_weight":5,"price":349,"sku":"abc-teste-123","height":0.2,"width":0.2,"length":0.2},{"category":7,"amount":2,"unitary_weight":4,"price":556,"sku":"abc-teste-527","height":0.4,"width":0.6,"length":0.15}]}'
```

- GET: /metrics?last_quotes={{?}}:

`last_quotes` is optional

```bash
curl --location 'localhost:3000/frete-rapido/api/v1/metrics?last_quotes=2'
```

## Running tests

To run the tests, you can use the following command:

```bash
go test ./... -cover -coverprofile=coverage.out
```

This command will generate a `coverage.out` file with the coverage report.

## Built With

- [Go](https://golang.org/)
- [MongoDB](https://www.mongodb.com/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Issue

Please, take a look in [#1](https://github.com/josemateuss/backend-challenge-frete-rapido/issues/1) issue.

## Author

- **[José Mateus](https://github.com/josemateuss)**
