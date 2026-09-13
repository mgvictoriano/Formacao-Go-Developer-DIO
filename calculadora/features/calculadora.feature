# language: pt

Funcionalidade: Calculadora
  Como uma pessoa usando a calculadora
  Eu quero somar, subtrair, multiplicar e dividir números
  Para obter o resultado das operações

  Cenário: Somar dois números positivos
    Dado que tenho os números 2 e 3
    Quando eu somo os números
    Então o resultado deve ser 5

  Cenário: Subtrair um número do outro
    Dado que tenho os números 10 e 4
    Quando eu subtraio os números
    Então o resultado deve ser 6

  Cenário: Multiplicar dois números
    Dado que tenho os números 3 e 4
    Quando eu multiplico os números
    Então o resultado deve ser 12

  Cenário: Dividir dois números
    Dado que tenho os números 10 e 2
    Quando eu divido os números
    Então o resultado deve ser 5

  Cenário: Dividir por zero deve dar erro
    Dado que tenho os números 10 e 0
    Quando eu divido os números
    Então deve ocorrer um erro de divisão por zero

