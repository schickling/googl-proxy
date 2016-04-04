FROM golang

ADD . /go/src/github.com/schickling/googl-proxy
WORKDIR /go/src/github.com/schickling/googl-proxy

RUN CGO_ENABLED=0 go build -o main -v -a -installsuffix nocgo .

EXPOSE 80
CMD ["./main"]
