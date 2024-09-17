FROM --platform=linux/amd64 golang:1.22-alpine3.20 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV go env -w GO111MODULE=on
RUN apk update --no-cache && apk add --no-cache tzdata && apk add git

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod tidy -compat=1.22.3 && go mod download
COPY . .
COPY src/etc /app/etc
COPY migrations /app/migrations
RUN go build -ldflags="-s -w" -o /app/bin src/main.go
RUN ls


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/America/Bogota /usr/share/zoneinfo/America/Bogota
ENV TZ America/Bogota

WORKDIR /app
COPY --from=builder /app/etc /app/etc
COPY --from=builder /app/bin /app/bin
COPY --from=builder /app/migrations /app/migrations

CMD ["./bin", "-f", "etc/definition.yaml", "-m", "migrations"]