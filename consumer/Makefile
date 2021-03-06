.DEFAULT_GOAL := help

ROOTDIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

CONSUMER_CONFIG ?= ${ROOTDIR}/configs/default.yaml
TAG ?= v0.1.3

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: run-dev
run-dev: ## Run api server in development environment
	CONSUMER_CONFIG=${CONSUMER_CONFIG} go run ${ROOTDIR}/cmd/main.go

##@ Deployment

.PHONY: image
image: ## Build docker image
	docker build \
		-f ${ROOTDIR}/deployments/Dockerfile \
		-t redshoore/istio-consumer:${TAG} \
		.

.PHONY: push
push: ## Push docker image to docker hub
	docker push redshoore/istio-consumer:${TAG}
