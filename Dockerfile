FROM golang:1.7.4-alpine

LABEL name="trinis-locations"

# Mix of https://blog.golang.org/docker
# and https://blog.tutum.co/2015/01/27/getting-started-with-golang-on-docker/
COPY ./src /go/src/app
COPY ./data /tmp
# RUN echo $GOPATH
RUN go get -d -v app
RUN go install -v app

EXPOSE 8080

CMD ["/go/bin/app", "/tmp/trinis_locations.csv"]
