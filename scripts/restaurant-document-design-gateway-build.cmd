pushd ..

go mod download

echo 0.0.1-local > build/version.txt

set GOARCH=amd64
set GOOS=linux
go build -o scripts/sut/bin/app ./cmd/document-design-gateway/main.go

echo set_by_ci > build/version.txt

popd
