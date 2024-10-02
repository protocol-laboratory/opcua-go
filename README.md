# opcua-go

![License](https://img.shields.io/badge/license-Apache2.0-green) ![Language](https://img.shields.io/badge/Language-Go-blue.svg) [![version](https://img.shields.io/github/v/tag/protocol-laboratory/opcua-go?label=release&color=blue)](https://github.com/protocol-laboratory/opcua-go/releases) [![Godoc](http://img.shields.io/badge/docs-go.dev-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/protocol-laboratory/opcua-go) [![codecov](https://codecov.io/gh/protocol-laboratory/opcua-go/branch/main/graph/badge.svg)](https://codecov.io/gh/protocol-laboratory/opcua-go)

Here’s a simplified explanation of the directory structure and a mermaid diagram to represent the package dependencies:

## Project Structure

```
examples/
|-- client/
|-- server/
opcua/
|-- ua/
|-- |-- enc/
```

These packages' dependency relationship is represented in the following diagram:

```mermaid
graph TD
    examples --> opcua
    opcua --> ua
    enc --> ua
```
