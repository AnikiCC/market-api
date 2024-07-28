FROM golang:1.22.2

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

CMD [ "migrate", "-database", "postgres://postgres:postgres@postgres:5432/DB?sslmode=disable", "-path", "/migrations", "up"]
