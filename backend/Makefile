DC = docker compose

up:
	${DC} up -d

down:
	${DC} down

docker_db_cli:
	${DC} exec -it db mysql -u root -p

docker_api_cli:
	${DC} exec -it api bash

docker_batch_cli:
	${DC} exec -it batch bash

deploy:
	gcloud builds submit --region=asia-east1