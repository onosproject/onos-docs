
#######
# This Makefile contains all targets related to the documentation
#######

DOCS_VERIFY_SKIP ?= true
DOCS_LINT_SKIP ?= true

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
	docker run  $(DOCKER_RUN_DOC_OPTS) $(ONOS_DOCS_BUILD_IMAGE) mkdocs serve

# Utilities Targets for each step
docs-image:
	docker build -t $(ONOS_DOCS_BUILD_IMAGE) -f docs.Dockerfile ./

docs-build: docs-clean docs-image
	docker run $(DOCKER_RUN_DOC_OPTS) $(ONOS_DOCS_BUILD_IMAGE) /bin/bash -c "./docker-entrypoint.sh && mkdocs build && chown -R $(shell id -u):$(shell id -g) ./site"


docs-verify: docs-build
	@if [ "$(DOCS_VERIFY_SKIP)" != "true" ]; then \
		docker build -t $(ONOS_DOCS_CHECK_IMAGE) -f check.Dockerfile ./; \
		docker run --rm -v $(CURDIR):/app $(ONOS_DOCS_CHECK_IMAGE) /verify.sh; \
	else \
		echo "DOCS_VERIFY_SKIP is true: no verification done."; \
	fi

docs-lint:
	@if [ "$(DOCS_LINT_SKIP)" != "true" ]; then \
		docker build -t $(ONOS_DOCS_CHECK_IMAGE) -f check.Dockerfile ./ && \
		docker run --rm -v $(CURDIR):/app $(ONOS_DOCS_CHECK_IMAGE) /lint.sh; \
	else \
		echo "DOCS_LINT_SKIP is true: no linting done."; \
	fi

docs-clean:
	rm -rf $(SITE_DIR) 
	rm -rf content/onos-*

.PHONY: all docs-verify docs docs-clean docs-build docs-lint
