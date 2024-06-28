# WARNING: `make` *requires* the use of tabs, not spaces, at the start of each command
include .env
export

.DEFAULT_GOAL := help
# probably based off the incantation from gist.github.com/prwhite/8168133
.PHONY: help
help:		## Help command
	@awk 'BEGIN {FS = ":.*##"; printf "\033[36m\033[0m"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

start:			## Start containers in dev mode
	docker compose -f docker-compose.yml up -d

stop:			## Stop containers in dev mode
	docker compose -f docker-compose.yml down

destroy:		## Delete all containers and volumes in dev mode
	docker compose -f docker-compose.yml down -v

