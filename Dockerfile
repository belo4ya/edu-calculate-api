FROM golang:1.23-alpine as builder

ARG TARGETOS
ARG TARGETARCH

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -tags=viper_bind_struct -o server ./cmd/server

FROM distroless/static:nonroot as prod

WORKDIR /

COPY --from=builder /app/server .

USER 65532:65532

ENTRYPOINT ["/server"]
