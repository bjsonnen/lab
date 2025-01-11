FROM alpine:latest
# Current size: 13MB

RUN apk update
# Current size: 18MB
RUN apk add git
# Current size: 37MB

RUN apk del git
# Current size: 37MB
RUN apk add git
# Current size: 56MB
