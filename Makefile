GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
MAKE=make

BUILD_PATH=build
ENTRY=./cmd/hextream
BINARY_NAME=hextream

LNX_BUILD=$(build)/$(BINARY_NAME)
WIN_BUILD=$(build)/$(BINARY_NAME).exe

export GO111MODULE=on

.PHONY: release

all: test build linux
build: 
	$(GOBUILD) -o $(BUILD_PATH)/$(BINARY_NAME) -v $(ENTRY)
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BUILD_PATH)/$(BINARY_NAME)
	rm -f $(BUILD_PATH)/$(BINARY_NAME)
	rm -f $(BUILD_PATH)/$(BINARY_NAME).exe
run:
	$(GORUN) $(ENTRY)

# Interacting with hextream cli
install:
	$(GOINSTALL) $(ENTRY)
cluster:
	$(GOINSTALL) $(ENTRY) && $(BINARY_NAME) cluster


# Cross compilation
linux: $(LNX_BUILD)
windows: $(WIN_BUILD)

$(LNX_BUILD):
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_PATH)/$(BINARY_NAME) -v $(ENTRY)
$(WIN_BUILD):
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_PATH)/$(BINARY_NAME).exe -v $(ENTRY)