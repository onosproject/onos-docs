FROM onosproject/golang-build:v0.6.8 AS build
ENV GO111MODULE=on
COPY . /go/src/github.com/onosproject/onos-docs
RUN cd /go/src/github.com/onosproject/onos-docs && GOFLAGS=-mod=vendor make build-docs-manager

FROM onosproject/onos-docs-base:latest

COPY --from=build /go/src/github.com/onosproject/onos-docs/build/_output/docs-manager /usr/local/bin/docs-manager

