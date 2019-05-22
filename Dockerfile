FROM golang:1.12 as builder
ADD . /app/github.com/sigmonsays/graphspace
ENV GOPATH=/app
EXPOSE 7001
RUN apt-get update && apt-get install -y graphviz
RUN go get github.com/sigmonsays/graphspace/...
RUN go install github.com/sigmonsays/graphspace/...
ADD docker/graphspace.yaml /app/graphspace.yaml
CMD /app/bin/graphspace -loglevel trace -config /app/graphspace.yaml
