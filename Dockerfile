FROM golang:1.7.4-alpine

COPY ./src /go/src/app
# RUN echo $GOPATH
RUN go get -d -v app
RUN go install -v app

EXPOSE 80

CMD ["/go/bin/app"]
