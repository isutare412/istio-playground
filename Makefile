.DEFAULT_GOAL := help

ROOTDIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

NAMESPACE ?= dev

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Deployment

.PHONY: helm-consumer
helm-consumer: ## Deploy consumer helm chart
	helm -n ${NAMESPACE} upgrade --install \
		isp-consumer \
		${ROOTDIR}/deployments/helm/isp-consumer

.PHONY: helm-api
helm-api: ## Deploy api server helm chart
	helm -n ${NAMESPACE} upgrade --install \
		isp-api \
		${ROOTDIR}/deployments/helm/isp-api

.PHONY: helm-user
helm-user: ## Deploy user server helm chart
	helm -n ${NAMESPACE} upgrade --install \
		isp-user \
		${ROOTDIR}/deployments/helm/isp-user
