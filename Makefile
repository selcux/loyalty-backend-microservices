.PHONY: protos build clean-images

protos:
	protoc -I protobuf/ protobuf/*.proto --go_out=. --go-grpc_out=.
	#protoc -I protobuf/ protobuf/*.proto --go_out=plugins=grpc:.

clean-images:
	@skaffold delete -f ./build/skaffold/skaffold.yaml
#	@echo "---------------- Cleaning dangling Docker images ----------------"
#	@docker images -f "dangling=true" -q | xargs --no-run-if-empty docker rmi -f

#BUILD_COMMAND = docker-compose -f build/loyalty-dev/docker-compose.yml build
BUILD_COMMAND = skaffold build -f ./build/skaffold/skaffold.yaml

build-images:
ifdef IMAGE
	$(BUILD_COMMAND) $(IMAGE)
else
	$(BUILD_COMMAND)
endif

build: build-images clean-images

#destroy: clean-images
#	@docker-compose -f build/loyalty-dev/docker-compose.yml down

run:
	@(./scripts/create-k8s-ns.sh hyperledger)
	@skaffold dev -f ./build/skaffold/skaffold.yaml
#	@docker-compose -f build/loyalty-dev/docker-compose.yml up

test:
	@go test -v --bench -json ./... --benchmem

#clean: clean-images destroy

# test network path needs to be added to $PATH

NETWORK_SH_PATH = $$(which network.sh)
TEST_NETWORK_PATH = $$(dirname $(NETWORK_SH_PATH))
test-network-up:
	echo $(NETWORK_SH_PATH)
	echo $(TEST_NETWORK_PATH)
	cd $(TEST_NETWORK_PATH) && \
	pwd && \
	network.sh up createChannel -ca -c mychannel -s couchdb

test-network-down:
	cd $(TEST_NETWORK_PATH) && \
	pwd && \
	sudo network.sh down

destroy-chaincodes:
	docker ps --filter name='loyalty-cc-*' --filter status='exited' -aq | xargs docker rm

#destroy-with-chaincodes: destroy destroy-chaincodes

