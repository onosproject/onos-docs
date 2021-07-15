FROM onosproject/onos-docs-manager:latest

ENV PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/root/.local/bin
COPY ./  /mkdocs
WORKDIR /mkdocs
VOLUME /mkdocs