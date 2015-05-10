FROM golang:1.4
WORKDIR /go
ENV PATH=$PATH:bin/:/go/api/bin
ENV GOPATH=$GOPATH:/go:/go/api:/go/gofetch
RUN go get github.com/tools/godep
RUN go get github.com/jinzhu/gorm
RUN go get github.com/ant0ine/go-json-rest/rest

WORKDIR /go/api/src/api/
ADD . /go/api

WORKDIR /go/api/src/
CMD go run main.go
