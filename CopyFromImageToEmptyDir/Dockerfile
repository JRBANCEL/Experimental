# Build
FROM golang:1.13-buster as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app

# Package
FROM alpine:3.7
COPY --from=build /go/bin/app /
COPY app.log /var/log/
CMD ["/app"]
