FROM golang:1.23.3-alpine

WORKDIR /workspace

ENV CGO_ENABLED=0

ENTRYPOINT ["go", "build", "." ]