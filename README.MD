### Executar o programa:
executar os passos:<br/>
- Abrir o terminal e digitar:
- **docker-compose up -d**
- **docker exec -it gokafka bash** que será o envio das mensagens
- Abrir outro terminal e digitar
- **docker exec -it gokafka bash** que será o consumidor das mensagens
- Abrir outro terminal que será o console do kafka e digitar:
- **docker exec -it gokafka_kafka_1 bash**
- Executar no terminal do tópico
<br/>
**go run cmd/producer/main.go** para enviar uma mensagem.<br/>
- Executar no terminal do consumidor<br/>
**go run cmd/consumer/main.go** para consumir uma mensagem.
