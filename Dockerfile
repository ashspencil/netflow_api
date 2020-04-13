FROM golang:1.13-alpine AS env
LABEL maintianer="wittmann <wittmann35312@gmail.com>"

ENV GOPATH "/go"
ENV PROJECT_PATH "$GOPATH/src/nmg/netflow"

RUN apk add --no-cache git make g++ && \
  go get -u github.com/golang/dep/cmd/dep

COPY . $PROJECT_PATH
RUN cd $PROJECT_PATH && \
  dep ensure && \
  make bin/api_server && \
  mv bin/api_server /tmp/api_server

FROM alpine:3.7
COPY --from=env /tmp/api_server /bin/api_server
ENTRYPOINT ["api_server"]
