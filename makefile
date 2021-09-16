ifeq ($(OS),Windows_NT)
	database := .\database
	migrations := .\database\migrations
	cli := .\cmd\cli
	main := .\cmd\server\main.go
else
	database := database/*.go
	migrations := database/migrations/*.go
	main := *.go
endif

full_migrate: drop_database migrate_init migrate seed

test:
	go run *.go -name=${name} -project=${project}

#
#migrate_init:
#	go run $(database) -action=init
#
#migrate:
#	go run $(database) -mode=migration -action=migrate
#
#migrate_create:
#	go run $(database) -mode=migration -action=create -name=${name}
#
#seed:
#	go run $(database) -mode=seed -action=migrate
#
#seed_create:
#	go run $(database) -mode=seed -action=create -name=${name}
#
#migrate_rollback:
#	go run $(migrations) rollback
#
#drop_database:
#	go run $(database) -action=dropDatabase
#
#create_model:
#	go run $(cli) -mode=model -action=create -name=${name} && goimports -w ./
#
#create_api:
#	go run $(cli) -mode=api -action=create -name=${name} && goimports -w ./
#
#create_service:
#	go run $(cli) -mode=service -action=create -name=${name} && goimports -w ./
