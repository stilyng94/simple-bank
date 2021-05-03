postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=${user} -e POSTGRES_PASSWORD=${password} -d postgres:13-alpine

createDb:
	docker exec -it postgres createdb --username=${username} --owner=${owner} ${dbName}

dropDb:
	docker exec -it postgres dropdb ${dbName}

createEntity:
	go run entgo.io/ent/cmd/ent init ${entity}

entgen:
	go generate ./ent

test:
	go test -test.v -cover ./tests/...

describeSchema:
	go run entgo.io/ent/cmd/ent describe ./ent/schema

server:
	go run main.go

.PHONY: test postgres createDb dropdb createEntity entgen describeSchema server