ATLAS = atlas

db/migrate:
	${ATLAS} migrate apply \
	--dir "file://ent/migrate/migrations" \
	--url "mysql://root:password@localhost:3306/flehmen"

db/status:
	${ATLAS} migrate status \
	--dir "file://ent/migrate/migrations" \
	--url "mysql://root:password@localhost:3306/flehmen"

ent/generate:
	go generate ./ent

db/seed:
	go run seeds/main.go