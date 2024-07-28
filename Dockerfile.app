FROM golang:1.22.2

WORKDIR /app # <--- нормальное название

COPY . .

RUN go mod download
RUN go build -o app ./cmd # <--- билд ./cmd

EXPOSE 8080
CMD ["./app"] # <--- запуск приложения
