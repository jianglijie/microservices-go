FROM alpine:latest

RUN mkdir /ticker && mkdir /ticker/log

WORKDIR /ticker

COPY services /ticker/services
COPY ./config/dev.toml /ticker/config.toml

CMD ["./services"]