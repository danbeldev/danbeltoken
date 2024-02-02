FROM golang:1.19 as builder
WORKDIR /app
ADD . /app/
RUN go build -o /danbeltoken main.go

EXPOSE 1212
CMD [ "/danbeltoken" ]