build:
	go build -o ./dist/bin/main main.go

run:
	gin --appPort 3000 --port 5000 run main.go