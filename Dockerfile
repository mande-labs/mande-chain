FROM golang:1.18-alpine AS build
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3 alpine-sdk build-base
WORKDIR /src/app/
RUN apk add $PACKAGES
COPY . .
RUN make install

FROM alpine
RUN apk add --no-cache bash
COPY --from=build /go/bin/mandeNode /usr/local/bin/mandeNode
ENTRYPOINT ["/usr/local/bin/mandeNode"]
