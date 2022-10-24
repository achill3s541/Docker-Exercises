FROM golang:1.19.2 as development
WORKDIR /app
COPY ./src/main.go /app/
CMD ["go", "run", "main.go"]

FROM development as builder
RUN go mod init array && go build

FROM alpine:3.16.2
COPY --from=builder /app/array .
CMD ["./array"]