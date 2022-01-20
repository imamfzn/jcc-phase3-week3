.PHONY: compile build run

IMAGE        = jcc/kubernetes/gosample
VERSION      = 0.1
ODIR        := _output


$(ODIR):
	@mkdir -p $(ODIR)

compile: $(ODIR)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o $(ODIR)/gosample main.go

build:
	docker build -t $(IMAGE):$(VERSION) .

run:
	go run main.go
