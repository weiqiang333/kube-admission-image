FROM golang:1.14.2-alpine3.11

COPY . /go/src/github.com/weiqiang333/kube-admission-image
WORKDIR /go/src/github.com/weiqiang333/kube-admission-image
RUN go build


FROM alpine:3.11
WORKDIR /app
RUN adduser -h /app -D web
COPY --from=0 /go/src/github.com/weiqiang333/kube-admission-image/kube-admission-image /app/

RUN chown -R web:web *
USER web
ENTRYPOINT ["./kube-admission-image"]
EXPOSE 8080
