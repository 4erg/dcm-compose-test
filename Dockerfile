FROM golang:1.23-alpine AS build
WORKDIR /src
COPY main.go .
RUN go mod init dcm-compose-demo >/dev/null 2>&1 || true
RUN go build -o /out/web main.go

FROM alpine:3.20
COPY --from=build /out/web /usr/local/bin/web
EXPOSE 8080
CMD ["web"]
