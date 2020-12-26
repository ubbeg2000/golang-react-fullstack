build:
	go build -o ./dist/bin/main main.go

run:
	gin run --appPort 3000 --port 5000 main.go