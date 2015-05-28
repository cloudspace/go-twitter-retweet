FROM ubuntu:latest

MAINTAINER Chris Moore <chris@cloudspace.com>

ADD ./ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ADD ./microservice.yml /microservice.yml

ADD ./go-twitter-retweet /go-twitter-retweet

CMD chmod a+x /go-twitter-retweet
