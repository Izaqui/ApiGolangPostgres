#Iniciar container postgres com Docker
sudo docker run -d --name "Nome da Base de Dados" -p 5432:5432 -e POSTGRES_PASSWORD="sua senha" postgres:13.5
#Listar containes
sudo docker ps
#Executar container
sudo docker exec -it "nome da base de dados" -U postgres
