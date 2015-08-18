
default:
	@godep go build
	@ls -ltrh

setup: 
	@echo Installing developer tooling, godep and reflex
	go get github.com/tools/godep
	go get github.com/cespare/reflex/...
	go get golang.org/x/tools/cmd/cover
	go get github.com/vektra/mockery/...

.goxc.ok:
	@echo Installing crossbuild tooling. This will take a while...
	go get github.com/laher/goxc
	goxc -t
	touch .goxc.ok

watch:
	@reflex -g '*.go' make test

test:
	@godep go test -coverprofile=c.out

coverage: test
	@godep go tool cover -html=c.out

bump:
	@goxc bump

release: .goxc.ok
	godep save
	goxc

mocks:
	@mockery -name Converter

brew_sha:
	@shasum -a 256 $(ver)/blade_$(ver)_darwin_amd64.zip
.PHONY: default test setup release watch coverage mocks bump brew_sha

