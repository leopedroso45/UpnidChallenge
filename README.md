**UpChallenge** [![Go Report Card](https://goreportcard.com/badge/github.com/leopedroso45/UpnidChallenge)](https://goreportcard.com/report/github.com/leopedroso45/UpnidChallenge)
----
API que avalia transações de e-commerce e devolve um score de 0 a 100 de risco, sendo 0 (sem indícios de fraude) e 100 (com máximo risco de fraude).

### Pré-requisitos para rodar a aplicação

  [Docker](https://docs.docker.com/)
  [Docker-compose](https://docs.docker.com/compose/install/)

### Instalação

  Após instalar e rodar o Docker na sua máquina você estará apto a rodar nossa API.

  Agora você pode clonar o repositório [UpnidChallenge](https://github.com/leopedroso45/UpnidChallenge) e abrir na sua IDE favorita. Para o desenvolvimento eu utilizei a [GoLand](https://www.jetbrains.com/pt-br/go/) da [JetBrains](https://www.jetbrains.com/pt-br/).

### Você está pronto

  Após clonar e abrir o projeto, você pode rodar o seguinte comando: 

  ```docker-compose up```

  Os testes unitários são executados durante a construção do container, se tudo ocorrer como esperado, a seguinte mensagem vai aparecer no seu terminal: 

  ```"Server running..."```

Caso ocorra algum problema, certifique-se que o Docker está rodando.

 **Sobre a API**
----

  **URL**

  /v1.0/transactions

  **Method**

  `POST`

  **Data Params**

No corpo da requisição é necessário o envio de um ou mais objetos de transação no formato JSON.

  *Note que é preciso utilizar os colchetes mesmo quando é apenas um objeto.*

Os objetos de transação são compostos pelos seguintes dados:
```json
[
  {
    "id": "5f20325488b6d415454025af",
    "value": "768.29",
    "paid_at": "2019-03-25 07:23:20",
    "ip_location": "SC/BR",
    "card_hold_name": "Ashlee Swanson",
    "customer": {
      "id": "5f203254077bb9e7bae8056f",
      "name": "Ashlee Swanson",
      "birth_date": "2005-04-18",
      "state": "SC/BR",
      "phone": "48 99999-9999"
    }
  }
]
```

  **Success Response**
  **Code:** 200 Ok <br/>
    **Content** `{ "id" : "5f20325488b6d415454025af","score" : "9" }`
    
  **Error Response**
  **Code** 400 BAD REQUEST <br/>
    **Content** `Error trying to read the request body, check if everything is correct.`
    
  **Sample Call**

```javascript
var settings = {
  "url": "localhost:8000/v1.0/transactions",
  "method": "POST",
  "timeout": 0,
  "headers": {
    "Content-Type": "application/json"
  },
  "data": JSON.stringify([{"id":"5f20325488b6d415454025af","value":"768.29","paid_at":"2019-03-25 07:23:20","ip_location":"SC/BR","card_hold_name":"Ashlee Swanson","customer":{"id":"5f203254077bb9e7bae8056f","name":"Ashlee Swanson","birth_date":"2005-04-18","state":"SC/BR","phone":"48 99999-9999"}}]),
  };

$.ajax(settings).done(function (response) {
  console.log(response);
});
```
