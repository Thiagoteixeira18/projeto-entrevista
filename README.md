# Projeto de Entrevista - Sistema de Regras para Opções de Carros

Este projeto implementa uma estrutura de dados para gerenciar regras de dependência e exclusão mútua entre opções de carros em um aplicativo de aluguel de carros. 
O objetivo é garantir que as opções sejam selecionadas de acordo com as regras estabelecidas.

## Estrutura do Projeto

O projeto está organizado nas seguintes pastas:

- Carros: Contém os arquivos relacionados à implementação das regras

  - carros.go: Implementação da estrutura de dados e métodos para adicionar dependências e conflitos

  - carros_test.go: Testes unitários para verificar as regras

- main.go: Ponto de entrada do aplicativo que utiliza a estrutura RuleSet para verificar as regras
