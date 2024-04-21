FROM golang:alpine AS builder
RUN apk add alpine-sdk

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main ./cmd/.

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main /
COPY --from=builder /app/storage/mysql/migrations/ /migrations/
COPY --from=builder /app/templates/ /templates/

ENTRYPOINT ["/main"]