REPO = github.com/jimma/blogmarkdownconverter
BUILD_PATH = $(REPO)/cmd/b2m
.PHONY: clean build help

.DEFAULT_GOAL := help
## clean                 Remove all generated build files.
clean:
	rm -rf b2m
## build                 Build this tool
build: 
	go build cmd/b2m/b2m.go
## install               Install this tool
install: 
	go install $(BUILD_PATH)

help : Makefile
	@sed -n 's/^##//p' $<
