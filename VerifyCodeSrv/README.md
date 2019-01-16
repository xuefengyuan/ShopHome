# VerifyCodeSrv Service

This is the VerifyCodeSrv service

Generated with

```
micro new ShopHome/VerifyCodeSrv --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.VerifyCodeSrv
- Type: srv
- Alias: VerifyCodeSrv

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
./VerifyCodeSrv-srv
```

Build a docker image
```
make docker
```