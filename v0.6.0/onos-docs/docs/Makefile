# This Makefile contains all targets related to building onos-docs website

ONOS_DOCS_BUILD_IMAGE ?= onos-docs
ONOS_DOCS_CHECK_IMAGE ?= $(ONOS_DOCS_BUILD_IMAGE)-check

SITE_DIR := $(CURDIR)/site

DOCKER_RUN_DOC_PORT := 8000
DOCKER_RUN_DOC_MOUNTS := -v $(CURDIR):/mkdocs 
DOCKER_RUN_DOC_OPTS := --rm $(DOCKER_RUN_DOC_MOUNTS) -p $(DOCKER_RUN_DOC_PORT):8000
MDL_DIRs := ./content/onos-config ./content/onos-topo ./content/onos-ztp ./content/onos-gui \
			./content/onos-ric ./content/ran-simulator \
            ./content/onos-cli/docs/setup.md ./content/developers ./content/onos-test


# Default: generates the documentation into $(SITE_DIR)
docs: docs-build 

# Writer Mode: build and serve docs on http://localhost:8000 with livereload
docs-serve: docs-build 
	docker run -it  $(DOCKER_RUN_DOC_OPTS) $(ONOS_DOCS_BUILD_IMAGE) /bin/sh -c "cd "site" && python3 -m http.server 8000"


# Build onos-docs related images
docs-image:
	docker build -t $(ONOS_DOCS_BUILD_IMAGE) -f docs.Dockerfile ./

docs-check:
	docker run --rm -v $(DOCKER_RUN_DOC_OPTS) $(ONOS_DOCS_BUILD_IMAGE) mdl $(MDL_DIRs) -s md.style.rb

docs-build: docs-clean docs-image
	docker run -it  $(DOCKER_RUN_DOC_OPTS) $(ONOS_DOCS_BUILD_IMAGE) docs-manager ./configs/versions.yml ./content/ ./mkdocs.yml ./configs/menu/onos-menu.js.gotmpl ./configs/menu/onos-menu.css

docs-clean:
	rm -rf $(SITE_DIR) 
	rm -rf content/master
	rm -rf content/onos-*
	rm -rf content/v*
	rm -rf content/theme
	rm -rf content/ran*
	rm -rf content/tla*
	rm -rf content/experimental


.PHONY: all docs docs-clean docs-build docs-lint
