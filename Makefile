GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get


all: test build

build:
	$(GOBUILD) ./cmd/gedcom

test:
	$(GOTEST) -v ./...

coverage:
	$(GOTEST) -cover ./...

clean:
	$(GOCLEAN)
	rm -f gedcom
