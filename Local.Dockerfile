FROM golang:latest

ENV CGO_ENABLED=1

RUN go get github.com/cespare/reflex

ENTRYPOINT ["reflex", "-c", "/testService/reflex.conf"]
