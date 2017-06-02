# cidr-test-service

A CORS-enabled API to determine whether the source IP address falls within given CIDR blocks

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/cidr-test-service
```

### Testing

``make test``