ENV := $(shell cat ../.last_used_env || echo "not-set")
KHULNASOFT_DEBUG := ""
-include ../.env.${ENV}

.PHONY: test
test:
	KHULNASOFT_DOMAIN=$(DOMAIN_NAME) KHULNASOFT_DEBUG=$(KHULNASOFT_DEBUG) KHULNASOFT_API_KEY=$(KHULNASOFT_API_KEY) ./test_khulnasoft.sh
