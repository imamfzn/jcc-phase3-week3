.PHONY: compile build run

ODIR      := _output
IMAGE     = jcc/kubernetes/gosample
VERSION   ?= $(shell git show -q --format=%h)

$(ODIR):
	@mkdir -p $(ODIR)

compile: $(ODIR)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $(ODIR)/gosample main.go

build:
	docker build -t $(IMAGE):$(VERSION) .

run:
	go run main.go
