# cidr-test-service

A CORS-enabled API to determine whether the source IP address falls within given CIDR blocks

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make build
$ CIDR_TEST_SERVICE_CIDRS="10.0.0.0/8 172.16.0.0/12 192.168.0.0/16" CIDR_TEST_SERVICE_LISTEN_ADDR_":8080" ./bin/cidr-test-service
```

### Testing

``make test``
