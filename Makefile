# Colors
COLOR_RESET=\033[0m
GREEN=\033[32m

.PHONY: help
help: # Show help for each of the Makefile recipes
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "${GREEN}$$(echo $$l | cut -f 1 -d':')${COLOR_RESET}:$$(echo $$l | cut -f 2- -d'#')\n"; done

test-local: # Run all tests
	go test ./days/...

build-docker: # Build docker image
	docker build -t aoc-2024 .

test-docker: # Run tests in docker container
	docker run aoc-2024

docker: build-docker test-docker

tidy: # Updates the dependencies and formats all .go files
	go mod tidy -v
	go fmt ./...

