# HomeWeb Service

This is the HomeWeb service

Generated with

```
micro new ShopHome/HomeWeb --namespace=go.micro --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.web.HomeWeb
- Type: web
- Alias: HomeWeb

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./HomeWeb-web
```

Build a docker image
```
make docker
```