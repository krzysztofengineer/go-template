watch:
	air -c .air.toml

build: templ
	go build -o ./tmp/main .

templ:
	templ generate