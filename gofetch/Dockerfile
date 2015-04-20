FROM golang:1.4
WORKDIR /go
RUN mkdir gofetch
RUN echo $GOPATH
ENV PATH=$PATH:bin/:/go/gofetch/bin
ENV GOPATH=$GOPATH:/go:/go/gofetch
RUN go get github.com/tools/godep
RUN go get github.com/mattes/migrate

WORKDIR /go/gofetch/src/fetcher/
ADD . /go/gofetch
RUN go get

WORKDIR /go/gofetch
