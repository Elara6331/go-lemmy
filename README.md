# Go-Lemmy

Go bindings to the [Lemmy](https://join-lemmy.org) API, automatically generated from Lemmy's openapi spec using the generator in [cmd/gen](cmd/gen).

### v4 Note

The latest version of this library uses the Lemmy v4 API. That API version hasn't been fully released yet so most Lemmy instances are still using the v3 API. If you want to continue using the v3 API for compatibility with older Lemmy instances, import `go.elara.ws/go-lemmy/v3`. Otherwise, import `go.elara.ws/go-lemmy/v4`. Note that these versions purely track the current Lemmy API version, so breaking changes can occur if the OpenAPI spec changes in such a way to cause them.

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
