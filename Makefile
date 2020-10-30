prophetbot:
	export `cat .env | xargs` && go run cmd/prophetbot/main.go