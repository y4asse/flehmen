ATLAS = atlas

db_migrate:
	${ATLAS} migrate apply \
	--dir "file://ent/migrate/migrations" \
	--url "mysql://root:password@localhost:3306/flehmen"

db_status:
	${ATLAS} migrate status \
	--dir "file://ent/migrate/migrations" \
	--url "mysql://root:password@localhost:3306/flehmen"

ent_generate:
	go generate ./ent