FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /gosrc
COPY ./src /gosrc

RUN go get -d -v
RUN CGO_ENABLED=0 go build -o /go/bin/lqip-api

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/lqip-api /go/bin/lqip-api

ENTRYPOINT ["/go/bin/lqip-api"]
CMD ["-p", "80"]