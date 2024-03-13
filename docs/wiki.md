## openapi

```sh
cargo init api
cd api
pnpm init
pnpm add -D @openapitools/openapi-generator-cli @stoplight/prism-cli
---
curl -sL http://127.0.0.1:4010
```

```sh
# mac
brew install java
sudo ln -sfn $(brew --prefix)/opt/openjdk/libexec/openjdk.jdk /Library/Java/JavaVirtualMachines/openjdk.jdk
```

## openapi-gen

```sh
# test
pnpm prism
pnpm test-gen
cd test
cargo build

# dev
pnpm gen
cd openapi
cargo build
```


