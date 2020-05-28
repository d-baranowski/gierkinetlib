export GIERKINET_DATA_REPLAY_PATH=$(CURDIR)/../test/data-replay

test:
	go test -short ./... | grep -v "no test files"

test-no-cache:
	go test -short -count=1 ./... | grep -v "no test files"

integration-test-entry:
	go test -count=1 -timeout 120s -parallel 1 ./sessions | grep -v "no test files"

integration-tests: integration-test-entry
