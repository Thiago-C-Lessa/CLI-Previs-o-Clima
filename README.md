# CLI-Previsao-Clima 🌦️

**WeatherCLI** é uma ferramenta de linha de comando (CLI) escrita em Go que retorna a previsão do tempo para um local especificado. Por padrão, exibe a previsão do tempo para o Rio de Janeiro. A aplicação utiliza a [WeatherAPI](https://www.weatherapi.com/) para obter as informações climáticas.

---

## 📋 Funcionalidades

- Recebe o nome de um local como parâmetro e retorna a previsão do tempo.
- Usa o Rio de Janeiro como local padrão se nenhum parâmetro for fornecido.
- Exibe informações como temperatura, condição do tempo e outros dados úteis.
- Ensina sobre o uso de **structs** e consumo de APIs em Go.

---

## 🛠️ Tecnologias

- **Linguagem:** Go
- **API:** WeatherAPI
- **Conceitos:** Structs, consumo de APIs REST, manipulação de JSON.

---

## 🚀 Como usar

### Pré-requisitos
- Go instalado (versão 1.18 ou superior).
- Uma chave de API válida da [WeatherAPI](https://www.weatherapi.com/).

### Configuração

1. Clone este repositório:
   ```bash
   git clone https://github.com/seuusuario/WeatherCLI.git
   cd WeatherCLI

## Melhorias 
- Mostrar apenas a previsão de horas futuras


![weatherT](https://github.com/user-attachments/assets/d52911f4-5634-43fa-b5e6-1542d2f08e5f)

### Colocando no terminal
Tive que retirar o uso do .envn então API_KEY virou variável global

![sem env](https://github.com/user-attachments/assets/1e714a1b-056a-4d3d-b1df-3ff4385890c2)
- Crie o executável com: "go build -o {nome-que-deseja} main.go"
- Transfira para /bin com: "sudo mv weathercli /usr/local/bin/"

### Saida sem input do usuário
![terminal](https://github.com/user-attachments/assets/814dfe88-b478-4ccd-b5b8-549946f8704a)

### Saida com input "London"
![terminal input](https://github.com/user-attachments/assets/0b9d8ea2-bc86-4712-aa87-8c8e477e084f)
