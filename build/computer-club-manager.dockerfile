FROM golang:1.21.8-alpine AS builder

ENV filepath ./input/test.txt

WORKDIR /build

ADD ../go.mod .

COPY .. .

RUN go build -o main ./cmd/main.go

FROM alpine:3

COPY --from=builder /build/main /
COPY --from=builder /build/${filepath} /

CMD ["sh", "-c", "./main ${filepath}"]