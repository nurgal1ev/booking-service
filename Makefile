DC = docker compose
DC_CONFIG = deployment/docker-compose.yml

up:
	$(DC) -f $(DC_CONFIG) up -d

.PHONY: down
down:
	$(DC) -f $(DC_CONFIG) down

.PHONY: restart
restart:
	$(DC) -f $(DC_CONFIG) restart

.PHONY: build
build:
	$(DC) -f $(DC_CONFIG) build

.PHONY: stop
stop:
	$(DC) -f $(DC_CONFIG) stop

.PHONY: logs
logs:
	$(DC) -f $(DC_CONFIG) logs --tail=100 app

.PHONY: ps
ps:
	$(DC) -f $(DC_CONFIG) ps

.PHONY: tables
tables:
	docker exec -it booking psql -U admin -d booking -c "\dt"