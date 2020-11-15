FROM golang:1.15-alpine

RUN mkdir /blog-service

ADD . /blog-service

WORKDIR /blog-service

RUN go build -o blog .

ENTRYPOINT [ "/blog-service/blog" ]

EXPOSE 50051
