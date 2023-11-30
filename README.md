# Workbench

## Description

This is my lil Go http server workbench.

It includes a lot of scaffolding for http servers that I've built up over the years including:

- api_common: a general purpose package with all of my util code for http routers, aws, caching, etc.
- swaggerui: a simple swagger ui page that can simplify documentation
- docker & air configs for deploying locally and enabling hot reloading
- a Makefile with a bunch of useful commands for local tinkering
- go-tuned pre-commit hooks for linting

## Repo Structure

- `api_common`: a general purpose package with all of my util code for http routers, aws, caching, etc.
- `handlers`: a package for all of the http handlers and the route registry
- `model`: a package for all of the data models used throughout the service
- `*-service`: an isolated package of business logic functions that can be deployed independently, or as a part of a larger service like `handlers`
- `swaggerui`: a simple swagger ui page that can simplify documentation based on yaml or json openapi files

## Usage

### Requirements

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)

### Commands

- `make`: run the server locally.  spins up mock services outlined in `docker-compose.yaml`
  - This starts the server on port 8080
  - Additional AWS infra can be connected to in `api_common/aws_utils/aws_client.go`

## Contributions

TODO (rsheikh) - complete this section & add contribution templates

## License
```text
Copyright 2023 Rachel Sheikh

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
