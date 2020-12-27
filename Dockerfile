FROM golang:latest

WORKDIR /app

COPY ./dist .

EXPOSE 3000:3000

CMD ["./dist/bin/main"]