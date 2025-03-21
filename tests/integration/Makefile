ENV := $(shell cat ../../.last_used_env || echo "not-set")
-include ../../.env.${ENV}

openapi := ../../spec/openapi.yml
codegen := go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: generate
generate:
	$(codegen) -old-config-style -generate client --package api $(openapi) > internal/api/client.gen.go
	$(codegen) -old-config-style -generate models --package api $(openapi) > internal/api/models.gen.go


.PHONY: build-debug
build-debug:
	go mod download
	go vet ./internal/...

@.PHONY: seed
seed:
	@echo "Applying seeds"
	@POSTGRES_CONNECTION_STRING=$(POSTGRES_CONNECTION_STRING) \
		TESTS_E2B_API_KEY=$(TESTS_E2B_API_KEY) \
		TESTS_SANDBOX_TEMPLATE_ID=$(TESTS_SANDBOX_TEMPLATE_ID) \
		TESTS_SANDBOX_BUILD_ID=$(TESTS_SANDBOX_BUILD_ID) \
		go run seed.go
	@echo "Done"

.PHONY: test
test:
	export TESTS_API_SERVER_URL=$(TESTS_API_SERVER_URL); \
	export TESTS_ORCHESTRATOR_HOST=$(TESTS_ORCHESTRATOR_HOST); \
	export TESTS_SANDBOX_TEMPLATE_ID=$(TESTS_SANDBOX_TEMPLATE_ID); \
	export TESTS_E2B_API_KEY=$(TESTS_E2B_API_KEY); \
	export TESTS_SUPABASE_TOKEN=$(TESTS_SUPABASE_TOKEN); \
	export TESTS_SANDBOX_TEAM_ID=$(TESTS_SANDBOX_TEAM_ID); \
	go test -v ./internal/main_test.go -count=1 && \
	go test -v ./internal/tests/... -count=1

.PHONY: connect-orchestrator
connect-orchestrator:
	CLIENT_IG=$$(gcloud compute instance-groups list \
		--filter="name~'^.*client.*'" \
		--format='value(name)' \
		--zones=$(GCP_ZONE) | head -n1) && \
	INSTANCE_ID=$$(gcloud compute instance-groups list-instances "$$CLIENT_IG" --zone=$(GCP_ZONE) --format='value(instance)' | head -n1) && \
	gcloud compute ssh "$$INSTANCE_ID" --zone=$(GCP_ZONE) -- -NL 5008:localhost:5008  -o PermitLocalCommand=yes -o LocalCommand="echo 'SSH tunnel established'"