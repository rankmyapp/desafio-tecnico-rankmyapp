FROM golang:1.24-alpine AS build
WORKDIR /app
RUN apk add --no-cache build-base
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/api

FROM alpine:3.19
RUN adduser -D -g '' appuser
WORKDIR /home/appuser
COPY --from=build --chmod=0755 --chown=appuser:appuser /out/api /usr/local/bin/api
USER appuser
EXPOSE 8080
CMD ["/usr/local/bin/api"]