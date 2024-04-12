Api completa em Golang, baseada no artigo 

    https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121

    Todos os créditos para o autor victor steven!

# Atualizar o Go no wsl (para versao 1.22.2):

## remover a versão atual do go:

    sudo rm -rf /usr/local/go* && sudo rm -rf /usr/local/go
    sudo rm /usr/bin/go

## instalar a nova versão

    VERSION=1.22.2
    OS=linux
    ARCH=amd64

    cd $HOME
    wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
    tar -xvf go1.22.2.linux-amd64.tar.gz
    mv go go-$VERSION
    sudo mv go-$VERSION /usr/local

## deixar configurada a versão automaticamente, quando abrir o wsl no windows

    cd ~
    nano .bashrc

    adicionar as seguintes linhas no final do arquivo, e salvar:

    export GOROOT=/usr/local/go-1.22.2
    export GOPATH=$HOME/projects/go
    export GOBIN=$GOPATH/bin
    export PATH=$PATH:$GOROOT/bin
    export PATH=$PATH:$HOME/projects/go/bin

## Pacotes: 

    go get github.com/badoux/checkmail
    go get github.com/jinzhu/gorm
    go get golang.org/x/crypto/bcrypt
    go get github.com/dgrijalva/jwt-go
    go get github.com/gorilla/mux
    go get github.com/jinzhu/gorm/dialects/mysql" //If using mysql 
    go get github.com/jinzhu/gorm/dialects/postgres //If using postgres
    go get github.com/joho/godotenv
    go get gopkg.in/go-playground/assert.v1

## Rodar o Postgres no docker:

    docker pull postgres
    docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=1234 postgres

## Rodar o app:

    go run main.go

## Rodar um teste:

    Exemplos: 

    go test -v --run TestUpdateAPost

    go test -v --run TestLogin

    go test -v --run TestCreateUser

    go test -v --run TestDeletePost

#Estrutura da API

    .env -> configurações da aplicação
    
    api
        auth        -> geração do token, validação do token
        controllers -> endpoints da aplicação e rotas (routes.go)
        middlewares -> middleware para controle de acesso da aplicação
        models      -> camada de negócio (contém a lógica da aplicação)
        responses   -> tratamento da resposta da API em json
        seed        -> migration com dados iniciais da aplicação
        util        -> rotinas genéricas da aplicação
        tests       -> testes unitários
        
