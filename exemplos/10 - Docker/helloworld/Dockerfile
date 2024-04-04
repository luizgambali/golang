# stage de build
FROM golang:1.22.1 AS build

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go

# stage imagem final
FROM scratch

WORKDIR /app

COPY --from=build /app/api ./

EXPOSE 8000

CMD [ "./api" ]