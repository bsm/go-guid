default: test

test:
	go test . -v 1

bench:
	go test -test.run=NONE -test.bench=. -test.benchmem

README.md: README.md.tpl $(wildcard *.go)
	becca -package $(subst $(GOPATH)/src/,,$(PWD))
