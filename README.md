# Go-Lemmy

Go bindings to the [Lemmy](https://join-lemmy.org) API, automatically generated from Lemmy's openapi spec using the generator in [cmd/gen](cmd/gen).

### Examples

Examples can be found in the [examples](examples) directory.

### How to generate

First, clone the [lemmy-js-client](https://github.com/LemmyNet/lemmy-js-client) repo at whatever version you need. For example:

```bash
git clone https://github.com/LemmyNet/lemmy-js-client -b 0.18.3
```

Inside it, build the openapi spec:

```bash
pnpm tsoa
```

Next, leave that directory and run:

```bash
go generate .
```

And that's it! Your generated code should be ready for use.
