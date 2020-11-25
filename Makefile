.DEFAULT_GOAL=test

DOCKER_TARGET_IMAGE := fluentd:1.0.0-local

.PHONY: run-env
run-env: fluentd-build elastic

.PHONY: clear
clear:
	find . -type f -name '*.out' -exec rm {} +
	find . -type f -name '*.test' -exec rm {} +

.PHONY: test
test:
	make -j 2 go-test fluentd-test

.PHONY: go-test
go-test:
	sleep 10s
	go test -v ./... -count 1
	@a=`docker ps | grep fluentd | awk '{ print $$1}'`; `docker stop $$a`

.PHONY: fluentd-test
fluentd-test:

	docker run -it --rm \
	--network docker_logs \
	-p "5000:5000" \
	-e "FLUENTD_CONF=/etc/fluent/fluent.conf" \
	-v "$(shell pwd)/fluent.conf:/etc/fluent/fluent.conf" \
	${DOCKER_TARGET_IMAGE} \
	/usr/bin/fluentd

.PHONY: elastic
elastic:
	docker-compose -f docker/docker-compose-elasticsearch.yaml up

.PHONY: fluentd-build
fluentd-build:
	@echo "clear image: " ${DOCKER_TARGET_IMAGE}
	-docker rm -f $(shell docker ps -a -q --filter=ancestor=${DOCKER_TARGET_IMAGE})
	-docker rmi -f $(shell docker images -q ${DOCKER_TARGET_IMAGE})
	-docker rmi $(shell docker images -f "dangling=true" -q)

	(cd docker/fluentd; docker build -t "${DOCKER_TARGET_IMAGE}" -f "$(shell pwd)/docker/fluentd/Dockerfile" .)