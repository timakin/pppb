pppb (Pretty Printer for Protocol Buffer)
=====

pppb is a data-inspector for protocol buffer serialized binary data.

## install

```
go get -u github.com/timakin/pppb
```

## usage

```
pppb inspect serialized.data -pkg=message.pkg.name -src=/path/to/proto_sources
```

## Contribution

1. Fork ([https://github.com/timakin/pppb/fork](https://github.com/timakin/pppb/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[timakin](https://github.com/timakin)
