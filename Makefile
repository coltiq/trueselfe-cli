.PHONY: clean build

clean:
	@rm -f data.db

build: clean
	@go build -o trueselfe main.go