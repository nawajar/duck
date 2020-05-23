FROM golang:1.14.3

ENV GO111MODULE=on

RUN mkdir -p /github.com/nawajar/duck/src

WORKDIR /github.com/nawajar/duck/src

COPY ./src .

RUN go mod download

RUN go get github.com/pilu/fresh

ENV PORT 8000 

CMD fresh -c runner.conf
