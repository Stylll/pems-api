#Commands

test:
	go test -cover ./tests

build:
	docker-compose build

run:
	docker-compose up
