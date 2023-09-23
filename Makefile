build:
	@go build -v -o go-pokemon-scraper .

run:
	@echo "RUN go-pokemon-scraper..."
	make build
	@./go-pokemon-scraper
