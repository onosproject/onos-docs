FROM alpine:3.9

ENV PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/root/.local/bin
RUN apk add --no-cache git
RUN apk add --no-cache bash
RUN apk add --no-cache openssh
COPY ./  /mkdocs
WORKDIR /mkdocs
VOLUME /mkdocs

RUN apk --no-cache --no-progress add py3-pip \
  && pip3 install --user -r requirements.txt
