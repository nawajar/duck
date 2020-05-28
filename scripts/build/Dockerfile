FROM golang:1.14.3
ENV GO111MODULE=on
RUN mkdir -p /github.com/nawajar/duck/
WORKDIR /github.com/nawajar/duck/
COPY ./src .
RUN go mod download
RUN go get github.com/pilu/fresh
RUN go build -o duck-api-app ./main.go

FROM alpine
WORKDIR /app
COPY --from=build-image /github.com/nawajar/duck/duck-api-app /app/
ENV PORT=8000 
ENTRYPOINT ./duck-api-app 
