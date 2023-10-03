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
