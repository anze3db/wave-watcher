FROM golang:1.4
WORKDIR /go
RUN mkdir api
ENV PATH=$PATH:bin/:/go/api/bin
ENV GOPATH=$GOPATH:/go:/go/api
RUN go get github.com/tools/godep
RUN go get github.com/jinzhu/gorm

WORKDIR /go/api/src/api/
ADD . /go/api
RUN go get

WORKDIR /go/api/src/
CMD go run main.go
