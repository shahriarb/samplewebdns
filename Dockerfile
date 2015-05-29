FROM debian:jessie

RUN apt-get update
RUN apt-get install -y dnsutils

ADD ./samplewebdns /app/samplewebdns

WORKDIR /app
