FROM golang:1.23.0-alpine AS base

RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd cmd
COPY internal internal

FROM base AS builder

WORKDIR cmd/perfin-api
RUN go build -o /build/perfin-api .

FROM base AS debug
WORKDIR cmd/perfin-api
RUN go build -gcflags="all=-N -l" -o /build/perfin-api .

ENV DEBUG_PORT=2345
EXPOSE $DEBUG_PORT
ENTRYPOINT /go/bin/dlv --listen=:$DEBUG_PORT --headless=true --api-version=2 --accept-multiclient exec /build/perfin-api

FROM alpine AS run
COPY --from=builder /build/perfin-api /
USER 1000
CMD ["/perfin-api"]

