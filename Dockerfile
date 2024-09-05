FROM golang:latest AS builder

WORKDIR /work
COPY . /work
RUN go build .

FROM debian:latest AS runner
WORKDIR /bin

COPY --from=builder /work/api /bin/api
RUN apt-get update && apt-get install -y \
    ca-certificates \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/bin/api", "--config", "/etc/api/config.toml"]
