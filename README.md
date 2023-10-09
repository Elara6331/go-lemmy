# Go-Lemmy

[![Go Reference](https://pkg.go.dev/badge/go.arsenm.dev/go-lemmy.svg)](https://pkg.go.dev/go.elara.ws/go-lemmy)

Go bindings to the [Lemmy](https://join-lemmy.org) API, automatically generated from Lemmy's source code using the generator in [cmd/gen](cmd/gen).

### Examples

Examples can be found in the [examples](examples) directory.

### How to generate

First, clone the [lemmy-js-client](https://github.com/LemmyNet/lemmy-js-client) repo at whatever version you need:

```bash
git clone https://github.com/LemmyNet/lemmy-js-client -b 0.18.3
```

Inside it, build the JSON docs file:

```bash
npm run docs -- --json docs.json
```

Next, build the generator:

```bash
go build ./cmd/gen
```

Remove all the existing generated code:

```bash
rm **/*.gen.go
```

Execute the generator:

```bash
./gen -json-file <path_to_docs.json> -out-dir .
```

And that's it! Your generated code should be ready for use.
