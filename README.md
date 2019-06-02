# Consulta CEP

### Realize consultas direto no portal dos [correios](http://www.buscacep.correios.com.br/sistemas/buscacep/BuscaCepEndereco.cfm)
Realize consultas em tempo real direto no portal dos correios.
O portal dos correios bloqueia quando é realizado uma quantidade em massa de requisições, utilize proxy para burlar esse bloqueio passando uma variavel de ambiente: CEP_PROXY=http://host:port  
##### Este projeto pode ser utilizado como biblioteca ou API Rest

#### Exemplo de Resposta
```json

[
   {
      "cep":"01310901",
      "logradouro":"Avenida Paulista, 1230 Edifício Sede BB SP Torre Matarazzo  Banco do Brasil S/A",
      "bairro":"Bela Vista",
      "municipio":"São Paulo",
      "uf":"SP"
   },
   {
      "cep":"01310932",
      "logradouro":"Avenida Paulista, 2202  Edifício Avenida Paulista",
      "bairro":"Bela Vista",
      "municipio":"São Paulo",
      "uf":"SP"
   }
]

```

## Usando imagem docker
```bash
docker run --name cep -p 8000:8000 -d casadosdados/consulta-cep
```
- É possível usar proxy, adicione ao comando `-e CEP_PROXY=http://host:port`

## Usando API REST
* URL: http://localhost:8000/consulta/cep
* Parametros:
    - q : Termo a ser pesquisado no portal
    - page : Número da página, entre 1 e 9
    - all : Captura todos os resultados de todas as páginas (Ignora o parametro page)
    
### Exemplo de requisição usando Curl
```bash
curl -H "Accept: application/json" "http://localhost:8000/consulta/cep?q=avenida paulista"
```

***
CSV com os CEPs capturados no portal dos correios (não existe nenhuma garantia que tenha todos os cep)
Este arquivo contém 888.596 ceps

https://github.com/casadosdados/consulta-cep/releases/download/0.0.2/cep-20190602.csv.gz