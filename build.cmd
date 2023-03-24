set GOOS=js
set GOARCH=wasm
go build -o ./site/dame.wasm .
cd site
python -m http.server