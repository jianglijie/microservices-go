FROM alpine:latest

RUN mkdir /ticker && mkdir /ticker/log

WORKDIR /ticker

COPY markets_service /ticker/services

CMD ["./services"]