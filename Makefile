setup:
	-brew install goreleaser/tap/goreleaser
	# hack to workaround go 1.11 modules, see 
	# https://github.com/vektra/mockery/issues/213
	# https://github.com/golang/go/issues/24250
	GO111MODULE=off go get gopkg.in/alecthomas/kingpin.v2
	GO111MODULE=off go get github.com/vektra/mockery/.../

mocks:
	mockery -all -dir pkg

test:
	go test ./pkg/...
deps:
	go mod tidy && go mod vendor

release:
	goreleaser --rm-dist

.PHONY: deps setup release mocks
