ifneq (,$(wildcard ./.env))
    include .env
    export
	ENV_FILE_PARAM = --env-file .env
endif

reset:
	@reset

air: reset
	@command -v air &>/dev/null || go install github.com/cosmtrek/air@latest
	@expect_unbuffer air -build.exclude_dir=tmp -c=.air.toml -d

app_test: reset
	@echo -n "\nstates: "
	@curl -sS http://localhost:${PORT}/states | jq '.[] | "\( .id ) - \( .label )"'
	@echo -n "\nusers: "
	@curl -sS http://localhost:${PORT}/users | jq '.[] | "\( .id ) - \( .name )"'
	@echo -n "\n\ntasks: "
	# @curl -sS http://localhost:${PORT}/tasks | jq '.[] | "\( .id ) - \( .id ) - \( .description )"'
	@curl -sS http://localhost:${PORT}/tasks | jq '.[]'
	@echo -n "\n"

debug:
	@dlv debug --listen=:2345 --headless=true --api-version=2 cmd/app/main.go

.PHONY: reset lint validate watch air modd app_main app_main_test
