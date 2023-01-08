# Go-Lemmy

[![Go Reference](https://pkg.go.dev/badge/go.arsenm.dev/go-lemmy.svg)](https://pkg.go.dev/go.arsenm.dev/go-lemmy)

Go bindings to the [Lemmy](https://join-lemmy.org) API, automatically generated directly from Lemmy's source code using the generator in [cmd/gen](cmd/gen).

Examples:

- HTTP: [examples/http](examples/http)
- WebSocket: [examples/websocket](examples/websocket)

### How to generate

First, build the generator:

```bash
go build ./cmd/gen
```

Clone Lemmy's source code at whatever version you need:

```bash
git clone https://github.com/LemmyNet/lemmy -b 0.16.7
```

Remove all the existing generated code:

```bash
rm **/*.gen.go
```

Execute the generator:

```bash
./gen -out-dir .
```

And that's it! Your generated code should be ready for use.