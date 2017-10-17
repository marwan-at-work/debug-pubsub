FROM golang:1.9.1

LABEL maintainer marwan.sulaiman@nytimes.com

RUN mkdir -p /go/src/github.com/marwan-at-work/debug-pubsub && \
    go get -u github.com/marwan-at-work/gowatch

COPY . /go/src/github.com/marwan-at-work/debug-pubsub

WORKDIR /go/src/github.com/marwan-at-work/debug-pubsub

CMD ["gowatch"]