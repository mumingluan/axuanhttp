# axuanhttp
no more axuan, so no more axuanhttp.
---
> HK4E HTTP Layer Emulators

## Build binaries

Setup your golang environment, then run the following commands to build binaries.

```bash
mkdir -p $GOPATH/src/github.com/mumingluan
cd $GOPATH/src/github.com/mumingluan
git clone https://github.com/mumingluan/axuanhttp.git
cd axuanhttp
go build -trimpath -ldflags "-s -w" -o bin/server main.go
# the binaries are in bin/ directory.
```

## Use Action artifacts

Go to Actions page and download artifacts for your OS and arch.
