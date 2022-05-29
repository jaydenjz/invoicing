.Phony: test
test:
	go test ./...

.Phony: compose-up
compose-up: ### Run docker-compose
	docker-compose up

.Phony: compose-down
compose-down: ### Run docker-compose
	docker-compose down --remove-orphans