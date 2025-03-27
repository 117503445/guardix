FROM alpine:3.21

WORKDIR /workspace

RUN apk add --no-cache nmap

COPY guardix guardix

ENTRYPOINT [ "./guardix" ]