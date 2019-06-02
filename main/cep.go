package main

import (
	"github.com/casadosdados/consulta-cep/correios"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/correios/cep", GETConsulta)
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
	if page <= 0 || page > 9 {
		return e.JSON(http.StatusNoContent, nil)
	}
	// correios não aceita consulta vazia
	if query == "" {
		return e.JSON(http.StatusNoContent, nil)
	}
	var result *correios.CollectionCEP
	// máximo de 1000 resultados
	if all != "" {
		result = correios.SearchALL(query)
		err = nil
	} else {
		result, err = correios.Search(query, page)
	}
	if err != nil{
		return e.JSON(http.StatusNotFound, map[string]interface{}{"error": err})
	}
	return e.JSON(http.StatusOK, result)


}
