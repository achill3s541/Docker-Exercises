FROM golang:1.19.2 as development
WORKDIR /app
COPY ./src/main.go /app/
CMD ["go", "run", "main.go"]

FROM development as builder
RUN go mod init array && go build

FROM alpine:3.16.2 as stage
WORKDIR /app
COPY --from=builder /app/array .
RUN adduser -D -H gouser && chown root:gouser -R . && chmod 750 -R .
USER gouser
CMD ["./array"]