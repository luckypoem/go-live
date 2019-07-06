FROM ubuntu:latest

WORKDIR /app

ADD ./main /app
ADD ./config.toml /app

EXPOSE 1935
EXPOSE 7001
EXPOSE 8090
EXPOSE 8040

ENTRYPOINT [ "/app/main" ]