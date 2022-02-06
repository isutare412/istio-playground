.DEFAULT_GOAL := help

ROOTDIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
ISTIO_OPERATOR_NS := istio-operator
ISTIO_SYSTEM_NS := istio-system
ISTIO_INGRESS_NS := istio-ingress

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

##@ Istio

.PHONY: istio-operator
istio-operator: ## Install istio operator
	kubectl create namespace ${ISTIO_OPERATOR_NS} --dry-run=client -o yaml | kubectl apply -f -
	helm -n ${ISTIO_OPERATOR_NS} upgrade --install \
		istio-operator \
		${ROOTDIR}/deployments/helm/infra-istio/istio-operator

.PHONY: istio
istio: ## Install istio
	kubectl create namespace ${ISTIO_SYSTEM_NS} --dry-run=client -o yaml | kubectl apply -f -
	kubectl create namespace ${ISTIO_INGRESS_NS} --dry-run=client -o yaml | kubectl apply -f -
	kubectl label namespace ${ISTIO_INGRESS_NS} istio-injection=enabled --overwrite
	kubectl -n ${ISTIO_SYSTEM_NS} apply -f ${ROOTDIR}/deployments/helm/infra-istio/istio

.PHONY: istio-addon
istio-addon: ## Install istio addons
	kubectl create namespace ${ISTIO_SYSTEM_NS} --dry-run=client -o yaml | kubectl apply -f -
	kubectl -n ${ISTIO_SYSTEM_NS} apply -f ${ROOTDIR}/deployments/helm/infra-istio/addon
