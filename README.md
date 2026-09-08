# Conversão de Escala Termométrica em GO

Um projeto prático que demonstra a sintaxe essencial do GO através da implementação de um algoritmo de conversão de temperaturas entre escalas diferentes.

## 📋 Descrição do Desafio

Este é um desafio proposto pela **[DIO - Digital Innovation One](https://www.dio.me/)** com o objetivo de:

- Explorar a sintaxe essencial do GO
- Criar um primeiro algoritmo prático
- Desafiar a lógica de programação
- Colocar em prática os comandos do core do GO

## 🎯 Objetivo

Desenvolver um programa em GO que realize conversão de temperaturas, especificamente convertendo a temperatura do ponto de ebulição da água de Kelvin para Celsius, aplicando conceitos fundamentais da linguagem.

## 🏗️ Estrutura do Projeto

```
conversao-escala-termometrica-exe-go/
├── README.md                    # Este arquivo
├── go.mod                       # Configuração do módulo GO
└── ebulição/
    └── main.go                  # Código principal do programa
```

## 📝 Conteúdo Principal

### `ebulição/main.go`

O arquivo contém:

- **Constante**: `pontoEbulicaoAguaKelvin` - temperatura em Kelvin (373.0 K)
- **Função**: `kelvinParaCelsius()` - converte temperatura de Kelvin para Celsius
- **Função**: `main()` - executa o programa e exibe os resultados formatados

**Fórmula utilizada:**
```
Celsius = Kelvin - 273
```

## 🚀 Como Executar

### Pré-requisitos

- GO 1.21 ou superior instalado
- Terminal/CMD com acesso ao seu ambiente

### Executando o programa

Na raiz do projeto, execute:

```bash
go run ebulição/main.go
```

Ou dentro do diretório `ebulição`:

```bash
cd ebulição
go run main.go
```

### Saída esperada

```
Ponto de ebulição da água
Kelvin:  373.0 K
Celsius: 100.0 °C
```

## 💡 Conceitos Aprendidos

- ✅ Declaração de pacotes (`package main`)
- ✅ Importação de bibliotecas (`import "fmt"`)
- ✅ Declaração de constantes (`const`)
- ✅ Definição de funções com parâmetros e retorno
- ✅ Uso da função `fmt.Printf()` para formatação de saída
- ✅ Operações aritméticas simples
- ✅ Tipos numéricos em GO (float64)


## 📚 Referências

- [Documentação oficial de GO](https://golang.org/doc/)
- [A Tour of Go](https://tour.golang.org/)
- [DIO - Digital Innovation One](https://www.dio.me/)

## 📄 Licença

Este projeto foi desenvolvido para fins educacionais.

---

**Desenvolvido com ❤️ como parte do programa de formação da DIO**
