FROM golang:latest
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app
# -d stop and not install
# -v verbose
RUN go get -d -v
RUN go install -v
CMD [ "app" ]
EXPOSE 8080
