package lemmy

//go:generate sed -i s/anyOf/oneOf/g lemmy-js-client/tsoa_build/swagger.json
//go:generate go tool ogen --package lemmy --target . --clean lemmy-js-client/tsoa_build/swagger.json
