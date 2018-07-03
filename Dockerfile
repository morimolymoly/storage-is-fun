FROM golang:latest

ADD . /go/src/github.com/morimolymoly/storage-is-fun

WORKDIR /go/src/github.com/morimolymoly/storage-is-fun
RUN go install github.com/morimolymoly/storage-is-fun/src/storage-is-fun
CMD ["/go/bin/storage-is-fun"]
