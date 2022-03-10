FROM golang:1.16-buster as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go build -o /go/bin/app

FROM scratch
COPY --from=build /go/bin/app /
CMD ["/app"]
