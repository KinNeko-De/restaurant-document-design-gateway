pushd ..

go mod download

set GOARCH=amd64
set GOOS=linux
go build -o scripts/sut/bin/app ./cmd/document-design-gateway/main.go

popd
