# Rodar Projeto

Install GO
- https://go.dev/doc/install

Rodar os comandos para inicializar os 3 projetos
- cd microservices/account && go run ./src/main.go start & .. && cd signin  && go run ./src/main.go start & .. && cd transactional && go run ./src/main.go start & .. && ..
- baixar o arquivo test.postman_collection.json e importar no postman com os endpoints e testar
- caso queira matar as portas
  - kill -9 $(lsof -t -i:28080 -sTCP:LISTEN) &&  kill -9 $(lsof -t -i:28081 -sTCP:LISTEN) &&  kill -9 $(lsof -t -i:28082 -sTCP:LISTEN)