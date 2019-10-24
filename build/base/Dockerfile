FROM alpine:3.9
RUN apk add libc6-compat
RUN apk add --no-cache git
RUN apk add --no-cache bash
RUN apk add --no-cache openssh

RUN apk --no-cache --no-progress add  py3-pip \
  && pip3 install --user mkdocs==1.0.4 \
  && pip3 install --user  pymdown-extensions==6.0 \
  && pip3 install --user mkdocs-bootswatch==1.0 \
  && pip3 install --user  mkdocs-material==4.0.2 \
  && pip3 install --user  markdown-include==0.5.1 \

