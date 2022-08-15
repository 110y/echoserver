FROM golang:1.19.0-bullseye AS builder

ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/github.com/110y/echoserver

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN make echocaller

## Runtime

FROM gcr.io/distroless/static:cd86c29059c6adf077464bc86a25c99b135a4f46
COPY --from=builder /go/src/github.com/110y/echoserver/bin/echocaller /usr/bin/echocaller

CMD ["/usr/bin/echocaller"]
