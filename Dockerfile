FROM golang:1.18-alpine AS build
RUN apk add curl make git libc-dev bash gcc linux-headers eudev-dev python3 alpine-sdk build-base

WORKDIR /src/app/
COPY . .
RUN make install

FROM alpine
RUN apk update
RUN apk add --no-cache bash
RUN apk add jq
COPY --from=build /go/bin/mandeNode /usr/local/bin/mandeNode
ENTRYPOINT ["/usr/local/bin/mandeNode"]
