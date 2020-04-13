VERSION_NUMBER ?= 0.0.2 #Remeber to modify version number by each time
VERSION ?= v$(VERSION_NUMBER)

GOOS ?= $(shell go env GOOS)

REPOPATH := netflow_api
BUILDTIME = $(shell date --rfc-3339=seconds)
COMMITID = $(shell git rev-parse HEAD)


$(shell mkdir -p ./bin)

.PHONY: build
build: bin/api_server

.PHONY: bin/api_server
bin/api_server:
	GOOS=$(GOOS) go build \
	  -ldflags="'-X version.buildTime=$(BUILDTIME) -X version.commitID=$(COMMITID)'" \
	  -a -o $@ cmd/main.go

.PHONY: dep
dep:
	dep ensure

.PHONY: build_image
build_image:
	docker build -t registry.gitlab.com/ashspencil2014/$(REPOPATH):$(VERSION) .

.PHONY: push_image
	docker push registry.gitlab.com/ashspencil2014/$(REPOPATH):$(VERSION)

.PHONY: run
run:
	./bin/server

.PHONY: clean
clean:
	rm -rf bin/
