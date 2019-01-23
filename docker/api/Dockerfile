FROM golang:1.10.2-alpine3.7

COPY src/api /go/src/api/

WORKDIR /go/src/api/

RUN /bin/sh -c /bin/sh -c apk update \
  && apk add --no-cache git \
  && go get -u github.com/jinzhu/gorm \
  && go get -u github.com/jinzhu/gorm \
  && go get -u github.com/go-sql-driver/mysql \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure

CMD gin -i run
