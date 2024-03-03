build:
	GOOS=linux GOARCH=amd64 go build -o build/linux/tools_scraping_py
	GOOS=darwin GOARCH=amd64 go build -o build/macos/tools_scraping_py