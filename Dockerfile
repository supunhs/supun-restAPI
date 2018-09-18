FROM golang:1.9

MAINTAINER Supun Karunarathna (supunkarunarathna92@gmail.com)

# install dependencies
RUN go get github.com/mongodb/mongo-go-driver/bson
RUN go get github.com/mongodb/mongo-go-driver/mongo
RUN go get github.com/mongodb/mongo-go-driver/mongo/findopt
RUN go get github.com/gorilla/mux

# env
ENV MONGO_HOST 192.168.1.102

# copy app
ADD . /app
WORKDIR /app

# build
RUN go build -o build/senz src/*.go

ENTRYPOINT ["/app/docker-entrypoint.sh"]