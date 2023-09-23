# ApiGolangPostgres

## Hello, I'm Izaquiel C. Silva,Studying Golang!👋

## Desenlvolvendo api com GoLang e postgres

# Comandos

## iniciando com a configuração do arquivo docker_run.sh

## Iniciar container postgres com Docker, devemos configurar o arquivo e depois excutar o script com o comando "./docker_run.sh" ou "sh docker_run.sh" ou "bash docker_run.sh"

### sudo docker run -d --name "Nome da Base de Dados" -p 5432:5432 -e POSTGRES_PASSWORD="sua senha" postgres:13.5
## Listar containes
### sudo docker ps
## Executar container
### sudo docker exec -it "nome da base de dados" -U postgres

# Logo após configure o arquivo config.toml

[api]
port = "9000"

[database]
host = "localhost"
#
user = ""
#
pass = ""
#
name = "api_Name"

# Iniciamos nosso server depois de configurar nosso arquivos e starta nosso banco.

# Execute "go run main.go"

## Nossa aplicação esta rodando.

<div style="display: inline_block"><br>
  <img align="center" alt="GoLang" height="30" width="40" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg">
  <img align="center" alt="Docker" height="30" width="40" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg">
</div>
  
##
 
<div> 
  <a href="https://www.instagram.com/izaquiel_silv/" target="_blank"><img src="https://img.shields.io/badge/-Instagram-%23E4405F?style=for-the-badge&logo=instagram&logoColor=white" target="_blank"></a>
  <a href = "mailto:izaquiel.silva.ads@gmail.com"><img src="https://img.shields.io/badge/-Gmail-%23333?style=for-the-badge&logo=gmail&logoColor=white" target="_blank"></a>
  <a href="https://www.linkedin.com/in/izaquiel-silva-dev/" target="_blank"><img src="https://img.shields.io/badge/-LinkedIn-%230077B5?style=for-the-badge&logo=linkedin&logoColor=white" target="_blank"></a> 
 
</div>
