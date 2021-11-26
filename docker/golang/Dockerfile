FROM golang:1.17-alpine

RUN apk update && \
    apk add \
    curl \
    gcc \
    git \
    musl-dev \
    && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

WORKDIR /go/src

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

CMD ["air"]