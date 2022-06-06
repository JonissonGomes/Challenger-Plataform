# Challenger-Plataform
<p>É um repositório criado para aprendizado de boas práticas, conhecimentos básicos de estrutura de dados, clean code, arquitetura e conexão com o banco. </p>

# Como executar o projeto localmente
Verifique se o docker-compose.yaml está na sua pasta e execute-o usando o comando:

```
docker-compose up -d
```
Após inicialize o GO MOD

```
go mod init "caminho_para_o_repositorio"`
```
# Dependências utilizadas

- BCrypt

Para instalar o BCrypt, basta utilizar o seguinte comando:
```
go get golang.org/x/crypto/bcrypt
```

- UUID

Para instalar o UUID, basta utilizar o seguinte comando:
```
go get github.com/satori/go.uuid
```

- GO RM
Para instalar o Go RM, basta utilizar o seguinte comando:
```
go get github.com/jinzhu/gorm
```

# Como instalar essas dependências
Utilizando o `go get "url_do_pacote"` você consegue estar instalando a dependências dentro do seu projeto 

Exemplo:
```
go get golang.org/x/crypto/bcrypt
```