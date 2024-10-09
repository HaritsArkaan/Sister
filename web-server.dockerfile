FROM golang:1.18-alpine as sister

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o webServerApp .

RUN chmod +x /app/webServerApp

FROM alpine:latest

RUN mkdir /app

COPY --from=sister /app/webServerApp /app

CMD [ "/app/webServerApp"]
