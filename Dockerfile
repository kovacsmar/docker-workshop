# Fetch
FROM dhi.io/golang:1-alpine3.23 AS fetch

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

# Generate
FROM ghcr.io/a-h/templ:v0.3.1001 AS generate

ENV GOMODCACHE=/go/pkg/mod

COPY --chown=65532:65532 --from=fetch /go/pkg/mod /go/pkg/mod
COPY --chown=65532:65532 . /build

WORKDIR /build

RUN ["templ", "generate"]

# Test
# ...

# Build
FROM dhi.io/golang:1-alpine3.23 AS build

ENV GOMODCACHE=/go/pkg/mod

COPY --chown=65532:65532 --from=fetch /go/pkg/mod /go/pkg/mod
COPY --chown=65532:65532 --from=generate /build /build

WORKDIR /build

ARG TARGETOS=linux
ARG TARGETARCH=amd64

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
  go build -trimpath \
  -ldflags="-s -w" \
  -o /build/main ./cmd/main.go

# Final
FROM dhi.io/alpine-base:3.23 AS final

WORKDIR /app

COPY --chown=65532:65532 --from=build /build/main /app/main

EXPOSE 8080

ENTRYPOINT ["/app/main"]
