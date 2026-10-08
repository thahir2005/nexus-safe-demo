FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY main.go .

RUN go mod init nexus-safe-demo && \
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
    go build -trimpath -ldflags="-s -w" -o /nexus-safe-demo main.go

FROM scratch

COPY --from=builder /nexus-safe-demo /nexus-safe-demo

EXPOSE 8000

ENTRYPOINT ["/nexus-safe-demo"]
