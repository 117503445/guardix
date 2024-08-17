set -e

GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build .
curl -T guardix.exe "http://192.168.100.135:7412/public-writable/guardix.exe"