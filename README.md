# password-generator-api

 RESTful API with Golang & Python that generates strong passwords

## Installation

To install the necessary dependencies, run:

`go mod download`

## Usage

To start the API, run:

`go run main.go`

By default, the API listens on port 8080.

To generate a password, send a GET request to
`/passwords/:length`, where `:length` is the desired length of the password. For example:

`GET /passwords/16`

The API will return a JSON object containing the generated password:

`
{
    "password": "3+QI0J0k3GB8brpo"
}
`
