FROM golang:1.19.5-alpine AS build

RUN apk update
RUN apk upgrade
RUN apk add --no-cache git make build-base

WORKDIR /tmp/unbar-api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Disable cgo
# RUN CGO_ENABLED=0 go build -o ./bin/unbar-api ./cmd/unbar/main.go
RUN go build -o ./bin/unbar-api ./cmd/unbar/main.go

FROM alpine:3.17.1

COPY --from=build /tmp/unbar-api/bin/unbar-api /app/unbar-api

EXPOSE 8080

CMD ["/app/unbar-api"]