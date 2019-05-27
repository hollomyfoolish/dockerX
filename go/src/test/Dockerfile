FROM golang:1.12.0
RUN mkdir -p /go/src/test
ADD ./main.go /go/src/test/
WORKDIR /go/src/test
ENTRYPOINT ["go", "run", "./main.go"]
