# Consulta CEP

### Realize consultas direto no portal dos [correios](http://www.buscacep.correios.com.br/sistemas/buscacep/BuscaCepEndereco.cfm)
### Este projeto pode ser utilizado como biblioteca ou API Rest

## Usando API REST
* URL: http://localhost:8000/consulta/cep
* Parametros:
    - q : Termo a ser pesquisado no portal
    - page : Número da página, entre 1 e 9
    - all : Captura todos os resultados de todas as páginas (Ignora o parametro page)
    
#### Exemplo de requisição usando Curl
```bash
curl -H "Accept: application/json" "http://localhost:8000/consulta/cep?q=avenida paulista"
```