.PHONY : build

export THIS := $(abspath $(lastword $(MAKEFILE_LIST)))
export THISD := $(dir $(THIS))

export GOPATH := $(THISD)

build: a b c d

a b c d:
	go build -v -o test_$@ github.com/gomatic/main_$@ && ./test_$@

