.PHONY: build clean-images

clean-images:
	@echo "---------------- Cleaning dangling Docker images ----------------"
	@docker images -f "dangling=true" -q | xargs --no-run-if-empty docker rmi -f

BUILD_COMMAND = docker-compose -f build/loyalty-dev/docker-compose.yml build

build-images:
ifdef IMAGE
	$(BUILD_COMMAND) $(IMAGE)
else
	$(BUILD_COMMAND) --parallel
endif

build: build-images clean-images

destroy: clean-images
	@docker-compose -f build/loyalty-dev/docker-compose.yml down

run:
	@docker-compose -f build/loyalty-dev/docker-compose.yml up

test:
	@go test -v --bench -json ./... --benchmem

clean: clean-images destroy