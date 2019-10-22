
#######
# This Makefile contains all targets related to the documentation
#######

DOCS_VERIFY_SKIP ?= true
DOCS_LINT_SKIP ?= false

ONOS_DOCS_BUILD_IMAGE ?= onos-docs
ONOS_DOCS_CHECK_IMAGE ?= $(ONOS_DOCS_BUILD_IMAGE)-check

SITE_DIR := $(CURDIR)/site

DOCKER_RUN_DOC_PORT := 8000
DOCKER_RUN_DOC_MOUNTS := -v $(CURDIR):/mkdocs 
DOCKER_RUN_DOC_OPTS := --rm $(DOCKER_RUN_DOC_MOUNTS) -p $(DOCKER_RUN_DOC_PORT):8000


# Default: generates the documentation into $(SITE_DIR)
docs: docs-clean docs-image docs-lint docs-build docs-verify

# Writer Mode: build and serve docs on http://localhost:8000 with livereload
docs-serve: docs-build
	docker run -it  $(DOCKER_RUN_DOC_OPTS) $(ONOS_DOCS_BUILD_IMAGE) /bin/sh -c "cd "site" && python3 -m http.server 8000"


# Utilities Targets for each step
docs-image:
	docker build -t $(ONOS_DOCS_BUILD_IMAGE) -f docs.Dockerfile ./

docs-build: docs-clean docs-image
	docker run -it  $(DOCKER_RUN_DOC_OPTS) $(ONOS_DOCS_BUILD_IMAGE) docs-manager ./configs/versions.yml ./content/ ./mkdocs.yml ./configs/menu/onos-menu.js.gotmpl ./configs/menu/onos-menu.css

docs-checker-image:
		docker build -t $(ONOS_DOCS_CHECK_IMAGE) -f check.Dockerfile ./

docs-verify: docs-checker-image docs-build
	@if [ "$(DOCS_VERIFY_SKIP)" != "true" ]; then \
		docker run --rm -v $(CURDIR):/app $(ONOS_DOCS_CHECK_IMAGE) /verify.sh; \
	else \
		echo "DOCS_VERIFY_SKIP is true: no verification done."; \
	fi

docs-clean:
	rm -rf $(SITE_DIR) 
	rm -rf content/master
	rm -rf content/onos-*
	rm -rf content/v*
	rm -rf content/theme
	rm -rf content/experimental


.PHONY: all docs-verify docs docs-clean docs-build docs-lint
