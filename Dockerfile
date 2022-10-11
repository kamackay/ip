FROM ubuntu:latest as go

ENV DEBIAN_FRONTEND noninteractive

WORKDIR /app/
WORKDIR $GOPATH/src/gitlab.com/kamackay/ip-lookup

RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y git apt-utils golang-go

ADD ./go.mod ./

RUN go mod download && \
    go mod verify

COPY ./ ./

RUN go build -tags=jsoniter -o application.file ./index.go && cp ./application.file /app/

FROM ubuntu:latest as stage

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y --reinstall apt-utils ca-certificates vim

WORKDIR /app

FROM scratch

COPY --from=stage / /

COPY --from=go /app/application.file /app/server

CMD /app/server
