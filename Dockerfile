FROM ubuntu:16.04

RUN apt-get update && apt-get install -y curl make git jq

RUN curl --insecure https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz | tar xz -C /usr/local

ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

RUN go get -u github.com/jteeuwen/go-bindata/...

RUN mkdir -p $GOPATH/src/github.com/ilackarms/
WORKDIR $GOPATH/src/github.com/ilackarms/unik

COPY ./ $GOPATH/src/github.com/ilackarms/unik

CMD make -e TARGET_OS=${TARGET_OS} localbuild && mv ./unik /opt/build/unik