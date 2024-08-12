**Rota:** `/financas/juroscomposto/simular`

**Método:** `GET`

**Parâmetros:**

* `quantidadeDeMeses`: Obrigatório. Número inteiro que representa a quantidade de meses a serem simulados.
* `valorInicial`: Obrigatório. Valor float64 que representa o valor inicial da aplicação financeira.
* `taxa`: Obrigatório. Valor float64 que representa a taxa de juros compostos.
* `valorAporte`: Obrigatório. Valor float64 que representa o valor do aporte inicial.
* `valorAumentoAporte`: Obrigatório. Valor float64 que representa o valor do aumento do aporte.
* `frequenciaAporte`: Opicional. String que representa a frequência do aporte (diário, mensal, etc.).
* `frequenciaAumentoAporte`: Opicional. String que representa a frequência do aumento do aporte (diário, mensal, etc.).

**Return:**

* Um objeto JSON com as seguintes propriedades:
	+ `aplicacao`: O resultado da simulação financeira.
	+ `mensagem`: Uma mensagem de status ("sucesso" ou "falha").

**Erros:**

* Se os parâmetros não forem válidos, o sistema retorna um erro JSON com a lista de erros e uma mensagem de status ("falha").
* Se houver alguma falha interna, o sistema retorna um erro JSON com uma mensagem de status ("falha").
