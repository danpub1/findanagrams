GOOS=js GOARCH=wasm go build -o findanagrams.wasm
gzip -9 -f -n findanagrams.wasm
go build -tags web -o findanagramsweb
go build
base64 findanagrams.wasm.gz | gawk '{print $0 "\\"}' >findanagrams.wasm.gz.b64
gawk -f embed.awk findanagrams.html.src >findanagrams.html
rm findanagrams.wasm.gz.b64