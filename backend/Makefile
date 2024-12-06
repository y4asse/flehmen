DC = docker compose

up:
	${DC} up -d

down:
	${DC} down

docker_db_cli:
	${DC} exec -it db mysql -u root -p -D flehmen

docker_db_clear:
	${DC} exec db mysql -u root -p -e "DROP DATABASE flehmen; CREATE DATABASE flehmen;"

docker_api_cli:
	${DC} exec -it api bash

docker_batch_cli:
	${DC} exec -it batch bash

db_diff:
	make -C app db_diff

db_migrate:
	make -C app db_migrate

db_status:
	make -C app db_status

db_seed:
	${DC} exec api make db_seed

deploy:
	gcloud builds submit --region=asia-east1