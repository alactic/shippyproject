# Shippyproject Service

This is the Shippyproject service

Generated with

```
micro new github.com/alactic/shippyproject --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.shippyproject
- Type: srv
- Alias: shippyproject

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

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
./shippyproject-srv
```

Build a docker image
```
make docker
```