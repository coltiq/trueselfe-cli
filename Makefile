.PHONY: clean build

clean:
	@rm -f data.db
	@rm -f .initialized

build: clean
	@go build -o trueselfe main.go