FROM golang:1.21.5 AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go install ./cmd/vaultic

FROM gcr.io/distroless/base-debian12

COPY --from=builder /go/bin/vaultic /usr/bin/vaultic

ENTRYPOINT ["/usr/bin/vaultic"]

