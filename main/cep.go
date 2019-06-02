package main

import (
	"github.com/casadosdados/consulta-cep/correios"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func main() {
	e := echo.New()
	e.GET("/consulta", GETConsulta)
	e.Logger.Fatal(e.Start(":8000"))
}

func GETConsulta(e echo.Context) error {
	query := e.QueryParam("q")
	pageQuery := e.QueryParam("page")
	all := e.QueryParam("all")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		page = 1
	}

	var result *correios.CollectionCEP
	if all != "" {
		result = correios.SearchALL(query)
		err = nil
	} else {
		result, err = correios.Search(query, page)
	}
	if err != nil{
		return e.JSON(http.StatusNotFound, err)
	}
	return e.JSON(http.StatusOK, result)


}
