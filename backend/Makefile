DC = docker compose

up:
	${DC} up -d

down:
	${DC} down

db/cli:
	${DC} exec -it db mysql -u root -p -D flehmen

db/clear:
	${DC} exec db mysql -u root -p -e "DROP DATABASE flehmen; CREATE DATABASE flehmen;"

api/cli:
	${DC} exec -it api bash

batch/cli:
	${DC} exec -it batch bash

agent/cli:
	${DC} exec -it agent bash

db/migrate:
	make -C app db/migrate

db/status:
	make -C app db/status

db/seed:
	${DC} exec api make db/seed

deploy:
	gcloud builds submit --region=asia-east1

logs/api:
	${DC} logs --follow api