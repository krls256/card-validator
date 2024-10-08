FROM golang:1.22-alpine3.18 AS builder


WORKDIR /app/
COPY . .

RUN go mod download
RUN set -x; apk add --no-cache \
    && CGO_ENABLED=0 go build -gcflags="all=-N -l"  \
    -a -installsuffix cgo -o ./bin/app cmd/server/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/bin .
COPY --from=builder /app/config.example.yml .
RUN cp config.example.yml config.yml

ENTRYPOINT ["./app"]