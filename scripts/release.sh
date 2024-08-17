set -e

GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build .
curl -T guardix.exe "http://100.64.0.5/public-writable/guardix.exe"