# Step 1: Modules caching
FROM golang:1.21.4-alpine3.17 as modules
COPY ../../go.mod ../../go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.21.4-alpine3.17 as builder
COPY --from=modules /go/pkg /go/pkg
COPY ../.. /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/main ./cmd/main.go

# Step 3: Final
FROM alpine
COPY --from=builder /app/config /app/config
COPY --from=builder /bin/main /app/main
WORKDIR /app
CMD ["/main"]