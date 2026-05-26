include .env
export

export PROJECT_ROOT=${shell pwd}

env-up:
	docker compose up -d todoapp-postgres

env-down:
	docker compose down todoapp-postgres

env-cleanup:
	@read -p "Are you sure you wanna clean up all env files? You might loose all the data. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres port-forwarder && \
		rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "files were cleaned"; \
	else \
		echo "Env cleanup was cancelled"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder


migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Enter parameter seq. Example: make migrate seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"


migrate-action:
	@if [ -z "$(action)" ]; then \
			echo "Enter parameter acction. Example: make migrate-action action=up"; \
			exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

migrate-up:
	@make migrate-action action=up
migrate-down:
	@make migrate-action action=down

logs-cleanup:
	@read -p "Are you sure you wanna delete all log files? You might loose all the data. [y/N]: " ans; \
    	if [ "$$ans" = "y" ]; then \
    		rm -rf ${PROJECT_ROOT}/out/logs && \
    		echo "Log files were cleaned"; \
    	else \
    		echo "Log files clean up was cancelled"; \
    	fi

todoapp-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/todoapp/main.go

todoapp-deploy:
	@docker compose up -d --build todoapp

todoapp-undeploy:
	@docker compose down todoapp

ps:
	@docker compose ps