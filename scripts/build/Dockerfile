FROM golang:1.14.3-alpine as build-image
RUN set -eux; \
        apk --no-cache add curl git; \
        curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh;

RUN set -eux; mkdir -p /github.com/nawajar/duck;        
ENV GO111MODULE=on
WORKDIR /github.com/nawajar/duck/
COPY ./src .
RUN go mod download
RUN go build -o duck-api-app ./main.go

FROM alpine
WORKDIR /app
COPY --from=build-image /github.com/nawajar/duck/duck-api-app /app/
ENV PORT=8000 
WORKDIR /app
ENTRYPOINT ./duck-api-app 

