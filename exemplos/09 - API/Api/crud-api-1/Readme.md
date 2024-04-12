Exemplo CRUD API, referencia: https://medium.com/@antonyshikubu/golang-crud-api-45abf75b6a10

Esta é um CRUD simples, que utiliza o framework Gin e o banco de dados Postgres

db.go       contém a inicialização da conexão com o banco de dados, struct e funções para o CRUD
router.go   contém as rotas


Para rodar o banco Postgresql com Docker:

docker pull postgres
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=1234 postgres

