FROM alpine:3.21

WORKDIR /workspace

RUN apk add --no-cache nmap curl

RUN curl -LO https://github.com/v-byte-cpu/sx/releases/download/v0.5.0/sx_0.5.0_linux_amd64.tar.gz && tar -xzf sx_0.5.0_linux_amd64.tar.gz && mv sx /usr/local/bin/sx && rm sx_0.5.0_linux_amd64.tar.gz

COPY guardix guardix

ENTRYPOINT [ "./guardix" ]