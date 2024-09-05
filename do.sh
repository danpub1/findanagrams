GOOS=js GOARCH=wasm go build -o findanagrams.wasm
gzip -9 -f -n findanagrams.wasm
go build -tags web -o findanagramsweb
go build