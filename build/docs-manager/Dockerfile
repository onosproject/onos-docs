# SPDX-FileCopyrightText: 2022 2020-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

FROM onosproject/golang-build:v1.3.0 AS build
ENV GO111MODULE=on
COPY . /go/src/github.com/onosproject/onos-docs
RUN cd /go/src/github.com/onosproject/onos-docs && GOFLAGS=-mod=vendor make build-docs-manager

FROM alpine:3.16
RUN apk add py-pip
RUN apk add python3-dev
RUN apk add libc6-compat
RUN apk add --no-cache --virtual .deps gcc musl-dev
RUN apk add --no-cache git bash openssh ruby && \
    gem install mdl -v 0.11.0

RUN apk --no-cache --no-progress add  python3 \
  && pip3 install --user mkdocs==1.1 \
  && pip3 install --user  pymdown-extensions==6.1 \
  && pip3 install --user  mkdocs-material==5.1.0 \
  && pip3 install --user  markdown-include==0.5.1 \
  && pip3 install --user mkdocs-minify-plugin

COPY --from=build /go/src/github.com/onosproject/onos-docs/build/_output/docs-manager /usr/local/bin/docs-manager

