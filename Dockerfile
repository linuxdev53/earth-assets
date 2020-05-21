
FROM golang:latest AS builder
RUN apt-get -y update && apt-get install -y ca-certificates

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/earth-assets

COPY go.mod .
RUN go mod download
COPY . .
RUN go install

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/earth-assets .
COPY --from=builder /go/src/earth-assets/assets ./assets
ENTRYPOINT ["./earth-assets"]
