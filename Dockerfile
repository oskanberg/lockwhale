ARG bin=/bin/lockwhale
ARG source=/go/src/github.com/oskanberg/lockwhale

FROM golang:1.11 as src

ARG source
COPY src ${source}

FROM src as build
ARG bin
ARG source

WORKDIR ${source}

RUN CGO_ENABLED=0 go build \
    -o ${bin} \
    -ldflags '-extldflags "-static"' \
    ./cmd/lockwhale

FROM scratch as run
ARG bin

COPY --from=build ${bin} /bin/service
ENTRYPOINT ["/bin/service"]
