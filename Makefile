export CGO_ENABLED=0
export GO111MODULE=on

DOCS_MANAGER_BUILD_VERSION := stable
DOCS_MANAGER_TEST_VERSION := latest

.PHONY: all docs docs-serve
docs: # @HELP Build documentation site
docs: clean deps build-docs-manager linters license_check images
	make -C ./docs docs
docs-serve: # @HELP Serve the documentation site localy.
docs-serve: deps build-docs-manager images
	make -C ./docs docs-serve

docs-serve: # @HELP Serve the documentation site localy.
docs-serve-without-build-image:
	make -C ./docs docs-serve-without-build-image

build-docs-manager: # @HELP build docs-manager application
	go build -o build/_output/docs-manager ./cmd/docs-manager


onos-docs-base-image:
	docker build . -f build/base/Dockerfile \
        		--build-arg DOCS_MANAGER_BUILD_VERSION=${DOCS_MANAGER_BUILD_VERSION} \
        		-t onosproject/onos-docs-base:${DOCS_MANAGER_TEST_VERSION}

onos-docs-manager-image:
	@go mod vendor
	docker build . -f build/docs-manager/Dockerfile \
    		--build-arg DOCS_MANAGER_BUILD_VERSION=${DOCS_MANAGER_BUILD_VERSION} \
    		-t onosproject/onos-docs-manager:${DOCS_MANAGER_TEST_VERSION}
	@rm -rf vendor


images: # @HELP build docs-manager application image
images: onos-docs-manager-image


publish: # @HELP publish version on github and dockerhub
	./../build-tools/publish-version ${VERSION} onosproject/onos-docs-manager

linters: # @HELP examines Go source code and reports coding problems
	golangci-lint run

license_check: # @HELP examine and ensure license headers exist        
	@if [ ! -d "../build-tools" ]; then cd .. && git clone https://github.com/onosproject/build-tools.git; fi
	./../build-tools/licensing/boilerplate.py -v --rootdir=${CURDIR}

deps: # @HELP ensure that the required dependencies are in place
	go build -v ./...
	bash -c "diff -u <(echo -n) <(git diff go.mod)"
	bash -c "diff -u <(echo -n) <(git diff go.sum)"


clean: # @HELP remove all the build artifacts
	make -C ./docs docs-clean
	rm -rf ./build/_output ./vendor

help:
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '
