build:
	npm --prefix ./frontend run build
	cp -r ./frontend/build ./dist/build
	go build -o ./dist/main main.go
	echo "Application built on dist"

compile:
	env go build -o dist/bin/main main.go
	env GOOS=windows GOARCH=386 go build -o dist/bin/main-windows-386 main.go
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/main-windows-amd64 main.go
	env GOOS=linux GOARCH=386 go build -o dist/bin/main-linux-386 main.go
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/main-linux-amd64 main.go

build-dev:
	npm --prefix ./frontend run build
	cp -r ./frontend/build/* ./static

run:
	go run main.go

dev-backend:
	gin --appPort 5000 --port 8000 run main.go

dev-frontend:
	npm --prefix ./frontend run start