# FROM alpine:3.21
# FROM registry.cn-hangzhou.aliyuncs.com/117503445-mirror/sync:linux.amd64.docker.io.library.alpine.3.21
FROM registry.cn-hangzhou.aliyuncs.com/117503445-mirror/sync@sha256:1c4eef651f65e2f7daee7ee785882ac164b02b78fb74503052a26dc061c90474

WORKDIR /workspace

RUN apk add --no-cache nmap curl

RUN curl -LO https://github.com/v-byte-cpu/sx/releases/download/v0.5.0/sx_0.5.0_linux_amd64.tar.gz && tar -xzf sx_0.5.0_linux_amd64.tar.gz && mv sx /usr/local/bin/sx && rm sx_0.5.0_linux_amd64.tar.gz

COPY guardix guardix

ENTRYPOINT [ "./guardix" ]