FROM golang:1.18-alpine AS build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go build -buildvcs=false --trimpath -o /go/bin/app

FROM scratch
COPY --from=build-env /go/bin/app /
CMD ["/app"]
