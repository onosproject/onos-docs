# SPDX-FileCopyrightText: 2022 2020-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

export CGO_ENABLED=0
export GO111MODULE=on

DOCS_MANAGER_BUILD_VERSION := stable
DOCS_MANAGER_TEST_VERSION := latest

build-tools:=$(shell if [ ! -d "./build/build-tools" ]; then cd build && git clone https://github.com/onosproject/build-tools.git; fi)
include ./build/build-tools/make/onf-common.mk

.PHONY: all docs docs-serve
docs: # @HELP Build documentation site
docs: clean deps build-docs-manager linters license images
	make -C ./docs docs

docs-serve: # @HELP Serve the documentation site localy.
docs-serve: deps build-docs-manager images
	make -C ./docs docs-serve

docs-serve-without-build-image: # @HELP Serve the documentation site localy.
docs-serve-without-build-image:
	make -C ./docs docs-serve-without-build-image

build-docs-manager: # @HELP build docs-manager application
	go build -o build/_output/docs-manager ./cmd/docs-manager

onos-docs-manager-image:
	@go mod vendor
	docker build . -f build/docs-manager/Dockerfile \
    		-t onosproject/onos-docs-manager:${DOCS_MANAGER_TEST_VERSION}
	@rm -rf vendor

images: # @HELP build docs-manager application image
images: onos-docs-manager-image

publish: # @HELP publish version on github and dockerhub
	./build/build-tools/publish-version ${VERSION} onosproject/onos-docs-manager

jenkins-test: jenkins-tools # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: deps build-docs-manager linters license images
	TEST_PACKAGES=NONE ./build/build-tools/build/jenkins/make-unit

jenkins-publish: # @HELP Jenkins calls this to publish artifacts
	./build/bin/push-images
	./build/build-tools/release-merge-commit
	./build/build-tools/build/docs/push-docs

clean:: # @HELP remove all the build artifacts
	make -C ./docs docs-clean
	rm -rf ./build/_output ./vendor
