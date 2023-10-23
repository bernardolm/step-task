ifneq (,$(wildcard ./.env))
    include .env
    export
	ENV_FILE_PARAM = --env-file .env
endif

clear:
	@clear

air:
	@command -v air &>/dev/null || go install github.com/cosmtrek/air@latest
	@expect_unbuffer air -build.exclude_dir=tmp -c=.air.toml -d

app_test: clear
	@echo -n "\n"
	@curl -sS http://localhost:${PORT}/users | jq '.'
	@echo -n "\n\n"
	@curl -sS http://localhost:${PORT}/tasks | jq '.'
	@echo -n "\n"


.PHONY: clear lint validate watch air modd app_main app_main_test
