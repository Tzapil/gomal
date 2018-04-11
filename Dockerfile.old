FROM golang:alpine
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app
# -d stop and not install
# -v verbose
RUN apk add --no-cache git mercurial \
    && go get -d -v \
    && apk del git mercurial
RUN go install -v
CMD [ "app" ]
EXPOSE 8080
