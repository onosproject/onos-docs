FROM alpine:3.9
RUN apk add libc6-compat
RUN apk add --no-cache git bash openssh ruby && \
    gem install --no-rdoc --no-ri mdl

RUN apk --no-cache --no-progress add  py3-pip \
  && pip3 install --user mkdocs==1.0.4 \
  && pip3 install --user  pymdown-extensions==6.0 \
  && pip3 install --user mkdocs-bootswatch==1.0 \
  && pip3 install --user  mkdocs-material==4.0.2 \
  && pip3 install --user  markdown-include==0.5.1 \
  && pip3 install --user mkdocs-minify-plugin \