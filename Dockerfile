FROM golang:alpine AS build

WORKDIR /app

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o main .

FROM scratch

WORKDIR /app

COPY --from=build /app/main /app/main

EXPOSE 4000

ENTRYPOINT [ "/app/main" ]