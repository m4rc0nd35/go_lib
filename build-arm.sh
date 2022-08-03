GOOS=linux GOARCH=arm CGO_ENABLED=1 GOARM="7" go build -buildmode=pie -ldflags "-s -w" -o test .
 #   go build -ldflags="-extldflags=-static" -o test.arm ./main.go