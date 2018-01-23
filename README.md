# P2P Network

[![Build Status](https://travis-ci.org/republicprotocol/go-network.svg?branch=master)](https://travis-ci.org/republicprotocol/go-network)
[![Coverage Status](https://coveralls.io/repos/github/republicprotocol/go-network/badge.svg?branch=master)](https://coveralls.io/github/republicprotocol/go-network?branch=master)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FMrTuoi%2Fgo-network.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FMrTuoi%2Fgo-network?ref=badge_shield)

> This library is a work in progress.

The P2P Network library is an official reference implementation of the P2P Network in the Republic Protocol, written in Go. It supports connecting nodes, and sending order fragments between nodes.

## Installation

There are several components that need to be installed before you can run tests.

### Install Proto3

Protobuf is a binary communication protocol developed by Google. The X Network uses it to perform remote procedure calls between miners. To install it, run the commands given below to install the required version of Protobuf.

```
curl -OL https://github.com/google/protobuf/releases/download/v3.2.0/protoc-3.2.0-linux-x86_64.zip
unzip protoc-3.2.0-linux-x86_64.zip -d protoc3

sudo mv protoc3/bin/* /usr/local/bin/
sudo mv protoc3/include/* /usr/local/include/

sudo chown $USER /usr/local/bin/protoc
sudo chown -R $USER /usr/local/include/google
```

The X Network is written in Go, so you will also need to install the Go plugin. If you are building an implementation in another language, you will need to install the Protobuf plugin for that language.

```
go get -u github.com/golang/protobuf/protoc-gen-go
```

### Install gRPC

gRPC is a remote procedure calling library developed by Google, built on top of Protobuf. Make sure you have followed the instructions for installing Protobuf, and then run the command below.

```
go get -u google.golang.org/grpc
```

## Tests

To run the test suite, install Ginkgo.

```
go get github.com/onsi/ginkgo/ginkgo
```

Now we can run the tests.

```
ginkgo -v --trace --cover --coverprofile coverprofile.out
```

## License

The P2P Network library was developed by the Republic Protocol team and is available under the MIT license. For more information, see our website https://republicprotocol.com.

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FMrTuoi%2Fgo-network.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FMrTuoi%2Fgo-network?ref=badge_large)