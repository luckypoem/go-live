FROM alpine:latest

WORKDIR /app

ADD ./main /app

EXPOSE 1935
EXPOSE 7001
EXPOSE 8090
EXPOSE 8040

RUN chmod 755 /app/main

ENTRYPOINT [ "./main" ]