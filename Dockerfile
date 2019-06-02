FROM golang:1.12 as builder
WORKDIR /correios
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /correios/cep main/cep.go

FROM scratch
LABEL maintainer="Casa dos Dados <contato@casadosdados.com.br> <https://casadosdados.com.br>"
LABEL version="0.0.1"

WORKDIR /correios
COPY --from=builder /correios/cep /correios/cep
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8000
ENTRYPOINT ["/correios/cep"]