##O QUE É UM MÓDULO?

Módulo em Go é uma forma de gerenciar dependências, listando e baixando todas as bibliotecas externas e componentes próprios no projeto.
Para tanto ele cria 2 arquivos, o `go.mod` (semelhante ao package.json ou ao Gemfile) e o `go.sum` (semelhante ao package-lock ou gemfile-lock).

Para usar bibliotecas de terceiros ou próprias sem ter de copiá-las no GOPATH é preciso usar o mod. O mod é uma evolução criada para não depender mais do GOPATH e conseguir rodar um projeto em outra máquina de forma mais fácil.

O arquivo `go.mod` mostra o nome do projeto (ou URL para baixá-lo se o nome do projeto for o repositório aonde ele será armazenado. Muito útil para encontrar a equipe encontrar o projeto e para bibliotecas), a versão do Go e a lista de bibliotecas externas (dependências) usadas, juntamente com suas versão. Ou seja, tem tudo que precisa para rodar o projeto.

O arquivo `go.sum` gera um checksum do mod para garantir que não houve alterações e firmando a lista de dependências e suas versões.

Vantagens de usar módulos:

- trabalhar fora do GOPATH e maior independência do projeto de configurações locais
- poder trabalhar com versões do código
- facilita gerenciamento de dependências e torna esse processo mais parecido com outras linguagens
- facilita reprodutibilidade dos builds

##VERSIONAMENTO DAS DEPENDÊNCIAS

Eu posso definir a versão que quero ao rodar o go get, adicionando `@version` (por ex: 1.12.3) ao final ou `@latest` para usar a mais atual. O go.mod e o go.sum guardarão a versão a ser usada evitando problemas de usar uam versão atual que não funcione com seu código (algo que o GOPATH não permitia).

O arquivo `go.mod` lista tanto suas dependências diretas (que você está usando no projeto) quanto as indiretas (usadas pelo projeto que você depende).

A biblioteca é baixada e instalada na sua máquina quando você roda o projeto (dar o `go run` ou roda a build).


##COMANDOS

`go mod init NOME_PROJETO` cria um arquivo go.mod com o nome do seu projeto (é uma boa prática o nome ser a URL do repositório)

`go get github.com/repositorio/projeto` adiciona uma biblioteca externa ao projeto para ser utilizada. Ele baixa a biblioteca dentro da pasta `GOPATH/pkg/mod` para ser usada como uma biblioteca padrão. O go get cria ou atualiza o arquivo go.sum.


`go get -u` checa se existem novas versões das bibliotecas usadas. Sempre bom usar para manter atualizado.

`go mod tidy` remove bibliotecas que não são mais utilizadas do go.mod e do go.sum. Basta remover elas do código que o comando entende que deve retirá-la do projeto. Também execute esse comando com frequência no projeto.

`go mod graph` mostra um grafo das dependências das bibliotecas que você usa.

`go mod vendor` cria uma pasta vendor para baixar e instalar as dependências dentro dele. Bom para não instalar as dependências no GOPATH e poluir ele e evitar conflitos na sua máquina. Isso também torna o projeto mais auto-contido.
