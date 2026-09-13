# Exercícios em Go

Repositório com exercícios desenvolvidos durante o estudo da linguagem Go, acompanhando minha formação em Go Developer.

Os exercícios passam por conceitos fundamentais da linguagem, estruturas de controle, funções, concorrência, API REST e testes automatizados com BDD.

## Projetos

| Projeto                                                            | O que exercita                        |
| ------------------------------------------------------------------ | ------------------------------------- |
| [Ponto de Ebulição da Água](#ponto-de-ebulição-da-água)            | Funções, constantes e tipos numéricos |
| [Divisíveis por 3](#divisíveis-por-3)                              | Laço `for` e operador módulo          |
| [Pin Pan](#pin-pan)                                                | `switch` e estruturas condicionais    |
| [Ping Pong](#ping-pong-concorrência)                               | Goroutines e canais (`chan`)          |
| [API REST — Cadastro de Clientes](#api-rest--cadastro-de-clientes) | HTTP, JSON, pacotes e `sync.Mutex`    |
| [Calculadora](#calculadora)                                        | Testes unitários e BDD com Gherkin    |

---

## Ponto de Ebulição da Água

Exercício que converte a temperatura de ebulição da água de Kelvin para graus Celsius, utilizando a fórmula:

```text
C = K - 273
```

📂 [`ebulição`](./ebulição)

Para executar:

```bash
go run ./ebulição
```

---

## Divisíveis por 3

Programa que percorre os números de 1 a 100 e exibe apenas aqueles que são divisíveis por 3.

A verificação utiliza o operador `%` (módulo) para verificar se o resto da divisão por 3 é igual a zero.

📂 [`problemas-numericos`](./problemas-numericos)

---

## Pin Pan

Variação do clássico FizzBuzz.

O programa percorre os números de 1 a 100 e:

* imprime `Pin` para múltiplos de 3;
* imprime `Pan` para múltiplos de 5;
* imprime `Pin Pan` para múltiplos de 3 e 5;
* imprime o próprio número nos demais casos.

A implementação utiliza `switch`, avaliando primeiro a condição mais específica — múltiplo de 15 — antes das condições mais genéricas.

📂 [`problemas-numericos`](./problemas-numericos)

---

## Ping Pong (concorrência)

Programa que imprime `ping` e `pong` alternadamente utilizando duas goroutines coordenadas por canais (`chan`).

Cada goroutine aguarda um sinal pelo canal antes de imprimir sua palavra e, em seguida, sinaliza a outra goroutine. Dessa forma, a alternância é controlada sem a utilização de locks (`Mutex`).

Um canal adicional é utilizado para sinalizar à função principal que a execução terminou.

📂 [`pingpong`](./pingpong)

Para executar:

```bash
go run ./pingpong
```

---

## API REST — Cadastro de Clientes

API REST desenvolvida em Go para cadastro de clientes de uma empreendedora do ramo de doceria.

O projeto utiliza o pacote [`gorilla/mux`](https://github.com/gorilla/mux) para roteamento HTTP e mantém os dados em memória utilizando um `map`, protegido por `sync.Mutex` para permitir acesso concorrente com segurança.

A API implementa as operações de CRUD:

| Método   | Rota             | Ação                            |
| -------- | ---------------- | ------------------------------- |
| `POST`   | `/clientes`      | Cadastra um novo cliente        |
| `GET`    | `/clientes`      | Lista todos os clientes         |
| `GET`    | `/clientes/{id}` | Busca um cliente pelo ID        |
| `PUT`    | `/clientes/{id}` | Atualiza os dados de um cliente |
| `DELETE` | `/clientes/{id}` | Remove um cliente               |

📂 [`api-rest`](./api-rest)

Para executar:

```bash
go run ./api-rest
```

---

## Calculadora

Implementação de uma calculadora com as quatro operações básicas:

* soma;
* subtração;
* multiplicação;
* divisão.

O projeto possui dois tipos de testes automatizados sobre o mesmo código:

1. **Testes unitários tradicionais**, utilizando o pacote de testes do Go;
2. **Testes BDD**, utilizando cenários escritos em Gherkin e executados com [Godog](https://github.com/cucumber/godog), uma implementação de Cucumber para Go.

Os cenários BDD utilizam a estrutura:

```text
Dado → Quando → Então
```

e estão escritos em português.

📂 [`calculadora/calculadora.go`](./calculadora/calculadora.go)
📂 [`calculadora/calculadora_test.go`](./calculadora/calculadora_test.go)
📂 [`calculadora/features/calculadora.feature`](./calculadora/features/calculadora.feature)

Para executar os testes:

```bash
go test ./calculadora/... -v
```

Para executar todos os testes do repositório:

```bash
go test ./... -v
```

---

## Tecnologias e conceitos estudados

* Go
* Funções
* Constantes
* Tipos numéricos
* Estruturas condicionais
* `for`
* `switch`
* Operador módulo (`%`)
* Goroutines
* Canais (`chan`)
* `sync.Mutex`
* HTTP
* JSON
* REST API
* Organização em pacotes
* Testes unitários
* BDD
* Gherkin
* Godog
* `go test`
