FROM golang:1.12
WORKDIR /correios
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /correios/cep main/cep.go

EXPOSE 8000

ENTRYPOINT ["/correios/cep"]