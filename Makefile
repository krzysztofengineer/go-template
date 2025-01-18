watch:
	air -c .air.toml

build: templ sqlc
	go build -o ./tmp/main .

templ:
	templ generate

sqlc:
	sqlc generate